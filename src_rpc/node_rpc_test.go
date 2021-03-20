package src_rpc

import (
	"chord_golang/pb"
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
	"testing"
)

func runServer(s *grpc.Server, server *ChordServer, done chan bool) {
	lis, err := net.Listen("tcp", server.Node.Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	pb.RegisterChordServer(s, server)
	done <- true
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func TestNodePing(t *testing.T) {

	done := make(chan bool)

	address_one := "127.0.0.1:10001"
	address_two := "127.0.0.1:10002"
	fake_address := "127.0.0.1:10003"

	//First Node
	Node_one := ChordNode{Id: 1,
					    Address: address_one}
	first_server := grpc.NewServer()

	//Second Node
	first_chord := ChordServer{Node: &Node_one}

	Node_two := ChordNode{Id: 2,
						  Address: address_two}

	//Servers
	second_server := grpc.NewServer()
	second_chord := ChordServer{Node: &Node_two}

	go runServer(first_server, &first_chord, done)
	go runServer(second_server, &second_chord, done)

	<-done
	<-done

	// Pinging first Node

	conn, err := grpc.Dial(Node_one.Address, grpc.WithInsecure())

	if err != nil {
		log.Printf("Could not contact Node %v", err)
	}

	client := pb.NewChordClient(conn)

	resp, err := client.Ping(context.Background(), &emptypb.Empty{})
	if err != nil {
		t.Fatalf("should have succeded %v", resp)
	}
	_ = conn.Close()
	// Pinging second Node

	conn, err = grpc.Dial(Node_two.Address, grpc.WithInsecure())

	if err != nil {
		log.Printf("Could not contact Node %v", err)
	}

	client = pb.NewChordClient(conn)

	resp, err = client.Ping(context.Background(), &emptypb.Empty{})
	if err != nil {
		t.Fatalf("should have succeded")
	}
	_ = conn.Close()
	// Pinging non existant

	conn, err = grpc.Dial(fake_address, grpc.WithInsecure())
	if err != nil {
		log.Printf("Could not contact Node %v", err)
	}
	client = pb.NewChordClient(conn)

	resp, err = client.Ping(context.Background(), &emptypb.Empty{})
	if err == nil {
		t.Fatalf("should have failed")
	}
	_ = conn.Close()

	first_server.Stop()
	second_server.Stop()

}

func TestFindSuccessorAndFindPredecessor(t *testing.T)  {

	// Initializing
	done := make(chan bool)
	var chordServers []*ChordServer
	var grpcServers [] *grpc.Server
	ids := []int32{5, 17, 22, 56, 71, 89, 92}
	var addr []string
	addrBase := 10000
	n := 7
	for i := 0; i < n; i++ {
		addr = append(addr, fmt.Sprintf("127.0.0.1:%v", i+addrBase))
	}
	for i := 0; i < n; i++ {
		chordServers = append(chordServers, &ChordServer{Node:
			&ChordNode{Address: addr[i], Id: ids[i]}})
		grpcServers = append(grpcServers, grpc.NewServer())
		go runServer(grpcServers[i], chordServers[i], done)
		//fmt.Println(*chordServers[i].Node)
	}
	for i := 0; i < n; i++ {
		<-done
	}

	// Manually Linking
	for idx, vals := range chordServers {

		if idx < len(chordServers)-1 {

			vals.SetSucc(context.Background(), &pb.Node{Id: chordServers[idx+1].Node.Id,
				Address: chordServers[idx+1].Node.Address})
		} else {
			vals.SetSucc(context.Background(), &pb.Node{Id: chordServers[0].Node.Id,
				Address: chordServers[0].Node.Address})
		}

	}

	chordServers[0].StabilizeAll(context.Background(), &empty.Empty{})

	for _, vals := range chordServers {

		fmt.Println(*vals.Node, *vals.Succ, *vals.SuccSucc)

	}

	// Find Successor tests

	base_case, _ := chordServers[0].FindSuccessor(context.Background(), &pb.FindSuccessorRequest{Id: 55})

	if base_case.Id != 56 {

		t.Errorf("Failed base case find successor %v", base_case.Id)

	}

	edge_one, _ := chordServers[6].FindSuccessor(context.Background(), &pb.FindSuccessorRequest{Id: 4})

	if edge_one.Id != 5 {

		t.Errorf("Failed edge case one find successor %v", edge_one.Id)

	}

	edge_two, _ := chordServers[6].FindSuccessor(context.Background(), &pb.FindSuccessorRequest{Id: 94})

	if edge_two.Id != 5 {

		t.Errorf("Failed edge case two find successor %v", edge_two.Id)

	}

	// Find Predecessor tests

	base_case, _ = chordServers[0].FindPredecessor(context.Background(), &pb.FindPredecessorRequest{Id: 91})

	if base_case.Id != 89 {

		t.Errorf("Failed base case find successor %v", base_case.Id)

	}

	edge_one, _ = chordServers[6].FindPredecessor(context.Background(), &pb.FindPredecessorRequest{Id: 4})

	if edge_one.Id != 92 {

		t.Errorf("Failed edge case one find successor %v", edge_one.Id)

	}

	edge_two, _ = chordServers[6].FindPredecessor(context.Background(), &pb.FindPredecessorRequest{Id: 16})

	if edge_two.Id != 5 {

		t.Errorf("Failed edge case two find successor %v", edge_two.Id)

	}

	// Stop all

	for _, val := range grpcServers{

		val.Stop()

	}

}

func TestJoinStabilize(t *testing.T) {

	done := make(chan bool)
	var chordServers []*ChordServer
	var grpcServers [] *grpc.Server
	ids := []int32{5, 17, 22, 56, 71, 89, 92}
	min := int32(1)
	max := int32(100)
	var addr []string
	addrBase := 10000
	n := 7
	for i := 0; i < n; i++ {
		addr = append(addr, fmt.Sprintf("127.0.0.1:%v", i+addrBase))
	}
	for i := 0; i < n; i++ {
		chordServers = append(chordServers, &ChordServer{Node:
		&ChordNode{Address: addr[i], Id: ids[i]}})
		grpcServers = append(grpcServers, grpc.NewServer())
		go runServer(grpcServers[i], chordServers[i], done)
	}
	for i := 0; i < n; i++ {
		<-done
	}

	// First Server Node
	fmt.Println("Inserting 5")
	chordServers[0].Minsize = min
	chordServers[0].Maxsize = max
	chordServers[0].SetSucc(context.Background(), &pb.Node{Id: chordServers[0].Node.Id,
		Address: chordServers[0].Node.Address})
	chordServers[0].SetSuccSucc(context.Background(), &pb.Node{Id: chordServers[0].Node.Id,
		Address: chordServers[0].Node.Address})
	chordServers[0].StabilizeAll(context.Background(), &empty.Empty{})

	// Nothing in it
	fmt.Println(*chordServers[0].Node, *chordServers[0].Succ, *chordServers[0].SuccSucc)
	fmt.Println("Inserting 92")
	// Joining 92 on 5
	chordServers[6].Minsize = min
	chordServers[6].Maxsize = max
	chordServers[0].Join(context.Background(), &pb.Node{Id: chordServers[6].Node.Id,
		Address: chordServers[6].Node.Address})
	chordServers[0].StabilizeAll(context.Background(), &empty.Empty{})

	if chordServers[6].Succ.Id != 5 {

		t.Errorf("Failed setting succ of 92, 5")

	}

	if chordServers[6].SuccSucc.Id != chordServers[6].Node.Id {

		t.Errorf("Failed setting succ succ of 92, 92")

	}

	if chordServers[0].Succ.Id != 92 {

		t.Errorf("Failed setting succ of 5, 92")

	}

	if chordServers[0].SuccSucc.Id != chordServers[0].Node.Id {

		t.Errorf("Failed setting succ succ of 5, 5")

	}


	fmt.Println(*chordServers[6].Node, *chordServers[6].Succ, *chordServers[6].SuccSucc)
	fmt.Println(*chordServers[0].Node, *chordServers[0].Succ, *chordServers[0].SuccSucc)

	fmt.Println("Inserting 17 and 56")
	// Joining 17 on 92 and 56 on 5
	chordServers[1].Minsize = min
	chordServers[1].Maxsize = max
	chordServers[3].Minsize = min
	chordServers[3].Maxsize = max

	chordServers[6].Join(context.Background(), &pb.Node{Id: chordServers[1].Node.Id,
		Address: chordServers[1].Node.Address})

	fmt.Println("Before stabilize")
	fmt.Println(*chordServers[0].Node, *chordServers[0].Succ)
	fmt.Println(*chordServers[1].Node, *chordServers[1].Succ)
	fmt.Println(*chordServers[6].Node, *chordServers[6].Succ)
	fmt.Println("----")

	chordServers[0].StabilizeAll(context.Background(), &empty.Empty{})
	fmt.Println("After stabilize")
	fmt.Println(*chordServers[0].Node, *chordServers[0].Succ, *chordServers[0].SuccSucc)
	fmt.Println(*chordServers[1].Node, *chordServers[1].Succ, *chordServers[1].SuccSucc)
	fmt.Println(*chordServers[6].Node, *chordServers[6].Succ, *chordServers[6].SuccSucc)
	fmt.Println("----")

	chordServers[0].Join(context.Background(), &pb.Node{Id: chordServers[3].Node.Id,
		Address: chordServers[3].Node.Address})

	chordServers[0].StabilizeAll(context.Background(), &empty.Empty{})

	//5->17->56-92->5
	//92->5->17
	//17->56-92
	//56->92->5

	fmt.Println(*chordServers[0].Node, *chordServers[0].Succ, *chordServers[0].SuccSucc)
	fmt.Println(*chordServers[1].Node, *chordServers[1].Succ, *chordServers[1].SuccSucc)
	fmt.Println(*chordServers[3].Node, *chordServers[3].Succ, *chordServers[3].SuccSucc)
	fmt.Println(*chordServers[6].Node, *chordServers[6].Succ, *chordServers[6].SuccSucc)

	if chordServers[6].Succ.Id != 5 {

		t.Errorf("Failed setting succ of 92, 5")

	}

	if chordServers[6].SuccSucc.Id != 17 {

		t.Errorf("Failed setting succ succ of 92, 17")

	}

	if chordServers[0].Succ.Id != 17 {

		t.Errorf("Failed setting succ of 5, 17")

	}

	if chordServers[0].SuccSucc.Id != 56 {

		t.Errorf("Failed setting succ succ of 5, 92")

	}


	for _, val := range grpcServers{

		val.Stop()

	}

}

func materializeAllNodes() ([]*ChordServer, [] *grpc.Server){

	ctx := context.Background()

	done := make(chan bool)
	var chordServers []*ChordServer
	var grpcServers [] *grpc.Server
	ids := []int32{5, 56, 22, 17, 89, 71, 92}
	min := int32(1)
	max := int32(100)
	addr := []string{"127.0.0.1:10000", "127.0.0.1:10001", "127.0.0.1:10002", "127.0.0.1:10003", "127.0.0.1:10004", "127.0.0.1:10005", "127.0.0.1:10006"}
	n := 7
	for i := 0; i < n; i++ {
		chordServers = append(chordServers, &ChordServer{Node:
		&ChordNode{Address: addr[i], Id: ids[i]}})
		grpcServers = append(grpcServers, grpc.NewServer())
		go runServer(grpcServers[i], chordServers[i], done)
		chordServers[i].Minsize = min
		chordServers[i].Maxsize = max
		chordServers[i].Shortcuts = make([]*ChordNode, 0)
	}
	for i := 0; i < n; i++ {
		<-done
	}
	chordServers[0].SetSucc(ctx, &pb.Node{Id: chordServers[0].Node.Id,
		Address: chordServers[0].Node.Address})
	chordServers[0].SetSuccSucc(ctx, &pb.Node{Id: chordServers[0].Node.Id,
		Address: chordServers[0].Node.Address})
	chordServers[0].StabilizeAll(context.Background(), &empty.Empty{})

	for i := 1; i < n; i++ {
		chordServers[0].Join(ctx, &pb.Node{Id: chordServers[i].Node.Id, Address: chordServers[i].Node.Address})
		chordServers[0].StabilizeAll(ctx, &empty.Empty{})
	}

	chordServers[0].AddShortcut(ctx, &pb.Node{Id: 56,
		Address: "127.0.0.1:10001"})
	chordServers[0].AddShortcut(ctx, &pb.Node{Id: 71,
		Address: "127.0.0.1:10005"})
	chordServers[2].AddShortcut(ctx, &pb.Node{Id: 89,
		Address: "127.0.0.1:10004"})

	return chordServers, grpcServers

}

/*

func TestClosestNodeTo(t *testing.T) {

	// Materializing non connected Nodes
	chordServers, grpcServers := materializeAllNodes()

	ctx := context.Background()

	// What's the closest Node, connected to 0, that would allow us to jump the closest to 18?, 17 or 22?
	closestToTestBase, _ := chordServers[0].ClosestNodeTo(ctx, &pb.ClosestNodeToRequest{Id: 18})
	closestToTestEdgeOne, _ := chordServers[0].ClosestNodeTo(ctx, &pb.ClosestNodeToRequest{Id: 16})
	closestToTestEdgeTwo, _ := chordServers[6].ClosestNodeTo(ctx, &pb.ClosestNodeToRequest{Id: 4})
	closestToTestEdgeThree, _ := chordServers[6].ClosestNodeTo(ctx, &pb.ClosestNodeToRequest{Id: 94})

	if closestToTestBase.Id != 22 {

		t.Errorf("ClosestNodeTo estimated wrong %v", closestToTestBase)

	}
	if closestToTestEdgeOne.Id != 17 {

		t.Errorf("ClosestNodeTo estimated wrong %v", closestToTestEdgeOne)

	}
	if closestToTestEdgeTwo.Id != 5 {

		t.Errorf("ClosestNodeTo estimated wrong %v", closestToTestEdgeTwo)

	}
	if closestToTestEdgeThree.Id != 5 {

		t.Errorf("ClosestNodeTo estimated wrong %v %v %v", closestToTestEdgeThree, chordServers[6].Minsize, chordServers[6].Maxsize)

	}


	for _, val := range grpcServers{

		val.Stop()

	}

}
 */

func TestShouldContainValue(t *testing.T) {
	// Based on the ring 5, 17, 22, 56, 71, 89, 92, 5
	// Base case in - 5 -> 17, looking for 16.
	// Base case not in - 5 -> 17, looking for 18
	// Edge case in - 92 -> 5, looking for 5

	if ShouldContainValue(17, 16, 5) != true {
		t.Errorf("cry")
	}
	if ShouldContainValue(17, 18, 5) != false {
		t.Errorf("cry")
	}
	if ShouldContainValue(5, 5, 92) != false {
		t.Errorf("cry")
	}
	if ShouldContainValueTwo(5, 5, 92) != true {
		t.Errorf("cry")
	}

}

func TestLookup(t *testing.T) {

	chordServers, grpcServers := materializeAllNodes()

	ctx := context.Background()

	lookupOne, _ := chordServers[0].Lookup(ctx, &pb.LookupRequest{Id: 101, Hops: 0})
	lookupTwo, _ := chordServers[1].Lookup(ctx, &pb.LookupRequest{Id: 4, Hops: 0})
	fmt.Println("Is it stopping here?")
	lookupThree, _ := chordServers[2].Lookup(ctx, &pb.LookupRequest{Id: 95, Hops: 0})
	fmt.Println("Or not?")
	lookupFour, _ := chordServers[3].Lookup(ctx, &pb.LookupRequest{Id: 32, Hops: 0})
	lookupFive, _ := chordServers[4].Lookup(ctx, &pb.LookupRequest{Id: 1, Hops: 0})
	lookupSix, _ := chordServers[5].Lookup(ctx, &pb.LookupRequest{Id: -1, Hops: 0})
	lookupSeven, _ := chordServers[6].Lookup(ctx, &pb.LookupRequest{Id: 16, Hops: 0})
	lookupEight, _ := chordServers[3].Lookup(ctx, &pb.LookupRequest{Id: 56, Hops: 0})

	// Testing Lookups

	if lookupOne != nil {
		t.Errorf("panik")
	}
	if lookupTwo.Node.Id != 5 {
		t.Errorf("panik")
	}
	if lookupThree.Node.Id != 5 {
		t.Errorf("panik")
	}
	if lookupFour.Node.Id != 56 {
		t.Errorf("panik")
	}
	if lookupFive.Node.Id != 5 {
		t.Errorf("panik")
	}
	if lookupSix != nil {
		t.Errorf("panik")
	}
	if lookupSeven.Node.Id != 17 {
		t.Errorf("panik")
	}
	if lookupEight.Node.Id != 56 {
		t.Errorf("panik")
	}

	// Testing Distances

	//fmt.Println(lookupOne.Hops)
	fmt.Println(lookupTwo.Hops)

	fmt.Println(lookupThree.Hops)

	fmt.Println(lookupFour.Hops)

	fmt.Println(lookupFive.Hops)
	//fmt.Println(lookupSix.Hops)
	fmt.Println(lookupSeven.Hops)
	fmt.Println(lookupEight.Hops)

	for _, val := range grpcServers{

		val.Stop()

	}

}

func TestLookupOnce(t *testing.T) {

	chordServers, grpcServers := materializeAllNodes()

	ctx := context.Background()

	for _, vals := range chordServers{

		fmt.Println(*vals.Node, *vals.Succ, *vals.SuccSucc)

	}

	lookupOne, _ := chordServers[0].Lookup(ctx, &pb.LookupRequest{Id: 87, Hops: 0})

	fmt.Println(lookupOne)

	for _, val := range grpcServers{

		val.Stop()

	}

}

func TestPinging(t *testing.T)  {

	chordServers, grpcServers := materializeAllNodes()

	ctx := context.Background()

	conn, err := grpc.Dial(chordServers[0].Node.Address, grpc.WithInsecure())

	if err != nil{

		t.Error("Panik")

	}

	c := pb.NewChordClient(conn)

	fmt.Println(c.Ping(ctx, &empty.Empty{}))

	_ = conn.Close()

	// Fake address

	conn, err = grpc.Dial("127.0.0.1:20000", grpc.WithInsecure())

	if err != nil{

		t.Error("Panik")

	}

	c = pb.NewChordClient(conn)

	ping, err := c.Ping(ctx, &empty.Empty{})

	if ping == nil {

		fmt.Println("I'm right")

	}

	_ = conn.Close()

	for _, val := range grpcServers{

		val.Stop()

	}

}

func TestMurder(t *testing.T) {

	chordServers, grpcServers := materializeAllNodes()
	fmt.Println("First Loop")
	for _, val := range chordServers {

		fmt.Println(val.Node.Id)

	}

	ctx := context.Background()

	fmt.Println("22 Shortcuts:", chordServers[2].Shortcuts)

	chordServers[4].Leave(ctx, &empty.Empty{})
	grpcServers[4].Stop()
	chordServers[0].StabilizeAll(ctx, &empty.Empty{})

	fmt.Println("22 Shortcuts:", chordServers[2].Shortcuts)

	fmt.Println("5 Shortcuts:", chordServers[0].Shortcuts)

	chordServers[1].Leave(ctx, &empty.Empty{})
	grpcServers[1].Stop()
	chordServers[0].StabilizeAll(ctx, &empty.Empty{})

	fmt.Println("5 Shortcuts:", chordServers[0].Shortcuts)

	chordServers[5].Leave(ctx, &empty.Empty{})
	grpcServers[5].Stop()
	chordServers[0].StabilizeAll(ctx, &empty.Empty{})

	fmt.Println("5 Shortcuts:", chordServers[0].Shortcuts)

	pos, _ := Find(chordServers, int32(0))

	var stabilizerNodeId int32
	for i, _ := range chordServers {

		if i != pos {

			stabilizerNodeId = int32(i)

		}

	}


	chordServers[0].Leave(ctx, &empty.Empty{})
	grpcServers[0].Stop()
	chordServers[stabilizerNodeId].StabilizeAll(ctx, &empty.Empty{})

	fmt.Println("Last Loop")
	for _, val := range chordServers {

		fmt.Println(val.Node.Id)

	}

	for idx, val := range chordServers {

		conn, _ := grpc.Dial(val.Node.Address, grpc.WithInsecure())

		c := pb.NewChordClient(conn)

		ping, _ := c.Ping(ctx, &empty.Empty{})

		fmt.Println("Pinging", val.Node.Id, "at",val.Node.Address)

		_ = conn.Close()

		if ping == nil {

			fmt.Println("Nil", val.Node.Id)

			chordServers[idx] = nil
			grpcServers[idx] = nil

		}

	}

	n := 0
	for idx, x := range chordServers {
		if x != nil {
			chordServers[n] = x
			grpcServers[n] = grpcServers[idx]
			n++
		}
	}
	chordServers = chordServers[:n]
	grpcServers = grpcServers[:n]

	for _, val := range chordServers {

		fmt.Println(val.Node.Id)

	}


	for _, val := range grpcServers{

		val.Stop()

	}

}

func TestNPlusOneJoin(t *testing.T) {

	chordServers, grpcServers := materializeAllNodes()
	for _, val := range chordServers {

		fmt.Println(val.Node.Id)

	}

	newNode := &ChordNode{
		Id:      int32(4),
		Address: fmt.Sprintf("127.0.0.1:%v", 10100+len(chordServers)),
	}


	done := make(chan bool)
	newChordServer := &ChordServer{Node: newNode}
	newGrpcServer := grpc.NewServer()
	go RunServer(newGrpcServer, newChordServer, done)
	<-done

	chordServers = append(chordServers, newChordServer)
	grpcServers = append(grpcServers, newGrpcServer)

	ctx := context.Background()

	fmt.Println(chordServers[0].HasValue(ctx, 92))

	/*

	chordServers[0].Join(ctx, &pb.Node{Id: chordServers[len(chordServers) - 1].Node.Id, Address: chordServers[len(chordServers) - 1].Node.Address})
	fmt.Println("Does it join?")

	for _, val := range chordServers{

		fmt.Println(val.Node.Id, "S-",val.Succ.Id, "NS-", val.SuccSucc.Id)

	}

	chordServers[0].StabilizeAll(ctx, &empty.Empty{})
	fmt.Println("After stabilize")

	for _, val := range chordServers{

		fmt.Println(val.Node.Id, "S-",val.Succ.Id, "NS-", val.SuccSucc.Id)

	}

	 */

	for _, val := range grpcServers{

		val.Stop()

	}

}

func TestLookupStartPredecessor(t *testing.T) {



}