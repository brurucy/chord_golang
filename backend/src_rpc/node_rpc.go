package src_rpc

import (
	"backend/pb"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type ChordNode struct{
	Id int32
	Address string
}

type ChordServer struct {
	node *ChordNode
	succ *ChordNode
	succSucc *ChordNode
	pred *ChordNode
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

func (s *ChordServer) Join(ctx context.Context, node *pb.Node) (*empty.Empty, error) {

	successor, _ := s.FindSuccessor(ctx, &pb.FindSuccessorRequest{Id: node.Id})
	predecessor, _ := s.FindPredecessor(ctx, &pb.FindPredecessorRequest{Id: node.Id})

	// Setting node's successor
	conn, err := grpc.Dial(node.Address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := pb.NewChordClient(conn)
	_, err = c.SetSucc(ctx, successor)
	if err != nil {
		return nil, err
	}
	// Setting node's predecessor
	conn, err = grpc.Dial(predecessor.Address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c = pb.NewChordClient(conn)
	_, err = c.SetSucc(ctx, node)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *ChordServer) FindSuccessor(ctx context.Context, request *pb.FindSuccessorRequest) (*pb.Node, error) {

	candidatePred := request.Id
	// Base case i.e seeking 70 from 17
	if candidatePred > s.node.Id && candidatePred < s.succ.Id {
		return &pb.Node{Id: s.succ.Id, Address: s.succ.Address}, nil
	// First Edge Case i.e seeking 4 from 92
	} else if candidatePred < s.node.Id && s.node.Id > s.succ.Id && candidatePred < s.succ.Id {
		return &pb.Node{Id: s.succ.Id, Address: s.succ.Address}, nil
	// Second Edge Case i.e seeking 94 from 92
	} else if candidatePred > s.node.Id && s.node.Id > s.succ.Id {
		return &pb.Node{Id: s.succ.Id, Address: s.succ.Address}, nil
	} else {
		node, err := s.succ.FindSuccessor(ctx, request.Id)
		if err != nil {
			return nil, err
		}
		return &pb.Node{Id: node.Id, Address: node.Address}, nil
	}

}

func (s *ChordServer) FindPredecessor(ctx context.Context, request *pb.FindPredecessorRequest) (*pb.Node, error) {

	candidatePred := request.Id
	// Base case i.e seeking 70 from 17
	if candidatePred > s.node.Id && candidatePred < s.succ.Id {
		return &pb.Node{Id: s.node.Id, Address: s.node.Address}, nil
		// First Edge Case i.e seeking 4 from 92
	} else if candidatePred < s.node.Id && s.node.Id > s.succ.Id && candidatePred < s.succ.Id {
		return &pb.Node{Id: s.node.Id, Address: s.node.Address}, nil
		// Second Edge Case i.e seeking 94 from 92
	} else if candidatePred > s.node.Id && s.node.Id > s.succ.Id {
		return &pb.Node{Id: s.node.Id, Address: s.node.Address}, nil
	} else {
		node, err := s.succ.FindPredecessor(ctx, request.Id)
		if err != nil {
			return nil, err
		}
		return &pb.Node{Id: node.Id, Address: node.Address}, nil
	}

}


func (n *ChordNode) FindSuccessor(ctx context.Context, id int32) (*ChordNode, error) {
	conn, err := grpc.Dial(n.Address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := pb.NewChordClient(conn)
	r, err := c.FindSuccessor(ctx, &pb.FindSuccessorRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return &ChordNode{r.Id, r.Address}, nil
}

func (n *ChordNode) FindPredecessor(ctx context.Context, id int32) (*ChordNode, error) {
	conn, err := grpc.Dial(n.Address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := pb.NewChordClient(conn)
	r, err := c.FindPredecessor(ctx, &pb.FindPredecessorRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return &ChordNode{r.Id, r.Address}, nil
}