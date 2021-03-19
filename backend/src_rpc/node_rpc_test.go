package src_rpc

import (
	"backend/pb"
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
	lis, err := net.Listen("tcp", server.node.Address)
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
	node_one := ChordNode{Id: 1,
					    Address: address_one}
	first_server := grpc.NewServer()

	//Second Node
	first_chord := ChordServer{node: &node_one}

	node_two := ChordNode{Id: 2,
						  Address: address_two}

	//Servers
	second_server := grpc.NewServer()
	second_chord := ChordServer{node: &node_two}

	go runServer(first_server, &first_chord, done)
	go runServer(second_server, &second_chord, done)

	<-done
	<-done

	// Pinging first node

	conn, err := grpc.Dial(node_one.Address, grpc.WithInsecure())

	if err != nil {
		log.Printf("Could not contact node %v", err)
	}

	client := pb.NewChordClient(conn)

	resp, err := client.Ping(context.Background(), &emptypb.Empty{})
	if err != nil {
		t.Fatalf("should have succeded %v", resp)
	}
	_ = conn.Close()
	// Pinging second node

	conn, err = grpc.Dial(node_two.Address, grpc.WithInsecure())

	if err != nil {
		log.Printf("Could not contact node %v", err)
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
		log.Printf("Could not contact node %v", err)
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
		chordServers = append(chordServers, &ChordServer{node:
			&ChordNode{Address: addr[i], Id: ids[i]}})
		grpcServers = append(grpcServers, grpc.NewServer())
		go runServer(grpcServers[i], chordServers[i], done)
		fmt.Println(*chordServers[i].node)
	}
	for i := 0; i < n; i++ {
		<-done
	}

	// Manually Linking
	for idx, vals := range chordServers {

		if idx < len(chordServers)-1 {

			vals.SetSucc(context.Background(), &pb.Node{Id: chordServers[idx+1].node.Id,
				Address: chordServers[idx+1].node.Address})
		} else {
			vals.SetSucc(context.Background(), &pb.Node{Id: chordServers[0].node.Id,
				Address: chordServers[0].node.Address})
		}
		fmt.Println(*vals.node, *vals.succ)
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
	var addr []string
	addrBase := 10000
	n := 7
	for i := 0; i < n; i++ {
		addr = append(addr, fmt.Sprintf("127.0.0.1:%v", i+addrBase))
	}
	for i := 0; i < n; i++ {
		chordServers = append(chordServers, &ChordServer{node:
		&ChordNode{Address: addr[i], Id: ids[i]}})
		grpcServers = append(grpcServers, grpc.NewServer())
		go runServer(grpcServers[i], chordServers[i], done)
	}
	for i := 0; i < n; i++ {
		<-done
	}

	// First Server Node
	fmt.Println("Inserting 5")
	chordServers[0].SetSucc(context.Background(), &pb.Node{Id: chordServers[0].node.Id,
		Address: chordServers[0].node.Address})
	chordServers[0].SetSuccSucc(context.Background(), &pb.Node{Id: chordServers[0].node.Id,
		Address: chordServers[0].node.Address})
	chordServers[0].StabilizeAll(context.Background(), &empty.Empty{})

	// Nothing in it
	fmt.Println(*chordServers[0].node, *chordServers[0].succ)
	fmt.Println("Inserting 92")
	// Joining 92 on 5
	chordServers[0].Join(context.Background(), &pb.Node{Id: chordServers[6].node.Id,
		Address: chordServers[6].node.Address})
	chordServers[0].StabilizeAll(context.Background(), &empty.Empty{})

	if chordServers[6].succ.Id != 5 {

		t.Errorf("Failed setting succ of 92, 5")

	}

	if chordServers[6].succSucc.Id != chordServers[6].node.Id {

		t.Errorf("Failed setting succ succ of 92, 92")

	}

	if chordServers[0].succ.Id != 92 {

		t.Errorf("Failed setting succ of 5, 92")

	}

	if chordServers[0].succSucc.Id != chordServers[0].node.Id {

		t.Errorf("Failed setting succ succ of 5, 5")

	}


	fmt.Println(*chordServers[6].node, *chordServers[6].succ, *chordServers[6].succSucc)
	fmt.Println(*chordServers[0].node, *chordServers[0].succ, *chordServers[0].succSucc)

	fmt.Println("Inserting 17")
	// Joining 17 on 92
	chordServers[6].Join(context.Background(), &pb.Node{Id: chordServers[1].node.Id,
		Address: chordServers[1].node.Address})
	chordServers[0].StabilizeAll(context.Background(), &empty.Empty{})

	fmt.Println(*chordServers[6].node, *chordServers[6].succ, *chordServers[6].succSucc)
	fmt.Println(*chordServers[0].node, *chordServers[0].succ, *chordServers[0].succSucc)
	fmt.Println(*chordServers[1].node, *chordServers[1].succ, *chordServers[1].succSucc)

	if chordServers[6].succ.Id != 5 {

		t.Errorf("Failed setting succ of 92, 5")

	}

	if chordServers[6].succSucc.Id != 17 {

		t.Errorf("Failed setting succ succ of 92, 17")

	}

	if chordServers[0].succ.Id != 17 {

		t.Errorf("Failed setting succ of 5, 17")

	}

	if chordServers[0].succSucc.Id != 92 {

		t.Errorf("Failed setting succ succ of 5, 92")

	}


	for _, val := range grpcServers{

		val.Stop()

	}

}