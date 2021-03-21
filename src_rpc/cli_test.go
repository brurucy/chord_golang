package src_rpc

import (
	"chord_golang/parser"
	"fmt"
	"google.golang.org/grpc"
	"testing"
)

func TestRead(t *testing.T) {

	var chordServers []*ChordServer
	var grpcServers []*grpc.Server

	file, _ := parser.Parse("../Input-file.txt")

	chordServers, grpcServers = Read(file.MinSize, file.MaxSize, file.Nodes, file.Shortcuts)

	for _, val := range chordServers {

		fmt.Println(val.Node.Id, val.Shortcuts, "S-", val.Succ.Id, "NS-", val.SuccSucc.Id)

	}

	for _, val := range grpcServers {

		val.Stop()

	}

}
