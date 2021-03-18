package src_rpc

import (
	"backend/pb"
	"context"
	"fmt"
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

func TestNodeJoin(t *testing.T)  {

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
	for _, val := range grpcServers{

		val.Stop()

	}


}