package src_rpc

import (
	"backend/pb"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"context"
	"log"
	"net"
)

func Find(slice []*ChordServer, val int32) (int, bool) {
	for i, item := range slice {
		if item.Node.Id == val {
			return i, true
		}
	}
	return -1, false
}

func RunServer(s *grpc.Server, server *ChordServer, done chan bool) {
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


func Materialize() ([]*ChordServer, []*grpc.Server) {

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
		go RunServer(grpcServers[i], chordServers[i], done)
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

func ShutDown(grpcServers []*grpc.Server) {


	for _, val := range grpcServers{

		val.Stop()

	}


}
