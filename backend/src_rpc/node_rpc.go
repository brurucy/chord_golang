package src_rpc

import (
	"backend/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

type ChordNode struct{
	Id int32
	Address string
}

type ChordServer struct {
	node *ChordNode
	succ *ChordNode
	succSucc *ChordNode
}

func (s *ChordServer) Ping(context.Context, *empty.Empty) (*pb.PingResponse, error) {

	var response pb.PingResponse

	response = pb.PingResponse{Alive: true}

	return &response,
		nil

}

func (s *ChordServer) SetSucc(ctx context.Context, node *pb.Node) (*empty.Empty, error) {
	s.succ = &ChordNode{
		Id:      node.Id,
		Address: node.Address,
	}
	return nil,nil
}

func (s *ChordServer) SetSuccSucc(ctx context.Context, node *pb.Node) (*empty.Empty, error) {
	s.succSucc = &ChordNode{
		Id:      node.Id,
		Address: node.Address,
	}
	return nil,nil
}
