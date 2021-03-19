package src_rpc

import (
	"backend/pb"
	"context"
	"fmt"
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
	return &empty.Empty{},nil
}

func (s *ChordServer) SetSuccSucc(ctx context.Context, node *pb.Node) (*empty.Empty, error) {
	s.succSucc = &ChordNode{
		Id:      node.Id,
		Address: node.Address,
	}
	return &empty.Empty{},nil
}

func (s *ChordServer) Stabilize(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {

	// Fixing succ
	//fmt.Println("Currently at", s.node.Id, " The Successor is: ", s.succ.Id)
	successor, _ := s.FindSuccessor(ctx, &pb.FindSuccessorRequest{Id: s.node.Id + 1})
	//fmt.Println("Getting succ", successor)
	s.succ = &ChordNode{Id: successor.Id, Address: successor.Address}
	// Fixing succ succ
	successorSuccessor, _ := s.FindSuccessor(ctx, &pb.FindSuccessorRequest{Id: s.succ.Id + 1})
	//fmt.Println("Getting succSucc", successorSuccessor)
	s.succSucc = &ChordNode{Id: successorSuccessor.Id, Address: successorSuccessor.Address}
	//fmt.Println("Currently at", s.node.Id, " The SuccSucc is: ", s.succSucc.Id)

	return &empty.Empty{}, nil
}

func (s *ChordServer) StabilizeAll(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {

	startId := s.node.Id
	// Pseudo do-while loop
	//fmt.Println("Stabilizing")
	_, err := s.Stabilize(ctx, e)
	if err != nil {
		fmt.Printf("Failed to Stabilize starting node: %v", err)
		return &empty.Empty{}, err
	}
	//fmt.Println("Passed the stabilize call")
	current := s.succ

	for current.Id != startId {

		fmt.Println("Currently at: ", current.Id)

		conn, err := grpc.Dial(current.Address, grpc.WithInsecure())
		if err != nil {
			fmt.Printf("Failed to Dial %v", err)
			return &empty.Empty{}, nil
		}
		c := pb.NewChordClient(conn)
		_, err = c.Stabilize(ctx, e)
		if err != nil {
			fmt.Printf("Failed to Stabilize %v", err)
			return &empty.Empty{}, nil
		}
		next, _ := c.FindSuccessor(ctx, &pb.FindSuccessorRequest{Id: current.Id + 1})

		_ = conn.Close()

		current = &ChordNode{Id: next.Id, Address: next.Address}
	}

	return &empty.Empty{}, nil

}

func (s *ChordServer) Join(ctx context.Context, node *pb.Node) (*empty.Empty, error) {

	successor, _ := s.FindSuccessor(ctx, &pb.FindSuccessorRequest{Id: node.Id})
	//log.Printf("Successor: %v", successor)
	predecessor, _ := s.FindPredecessor(ctx, &pb.FindPredecessorRequest{Id: node.Id})
	//log.Printf("Predecessor: %v", predecessor)
	// Updating Successors
	// Setting node's successor
	conn, err := grpc.Dial(node.Address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Error accessing the predecessor's server")
		return &empty.Empty{}, err
	}
	defer conn.Close()
	c := pb.NewChordClient(conn)
	_, err = c.SetSucc(ctx, successor)
	if err != nil {
		fmt.Println("Error setting the node's successor")
		return &empty.Empty{}, err
	}
	// Setting the node's predecessor successor
	conn, err = grpc.Dial(predecessor.Address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Error accessing the predecessor's server")
		return &empty.Empty{}, err
	}
	defer conn.Close()
	c = pb.NewChordClient(conn)
	_, err = c.SetSucc(ctx, node)
	if err != nil {
		fmt.Println("Error setting predecessor's successor")
		return &empty.Empty{}, err
	}
	return &empty.Empty{}, err
}

func (s *ChordServer) FindSuccessor(ctx context.Context, request *pb.FindSuccessorRequest) (*pb.Node, error) {

	candidatePred := request.Id
	// Base case i.e seeking 70 from 17
	if candidatePred > s.node.Id && candidatePred < s.succ.Id {
		return &pb.Node{Id: s.succ.Id, Address: s.succ.Address}, nil
	// First Edge Case i.e seeking 4 from 92
	} else if candidatePred < s.node.Id && s.node.Id > s.succ.Id && candidatePred < s.succ.Id {
		return &pb.Node{Id: s.succ.Id, Address: s.succ.Address}, nil
	// Second Edge Case i.e seeking 94 from 92, where 92 is the LAST node before the FIRST
	} else if candidatePred > s.node.Id && s.node.Id > s.succ.Id {
		return &pb.Node{Id: s.succ.Id, Address: s.succ.Address}, nil
	// Third Edge Case i.e, where a node equals its successor
	} else if candidatePred > s.node.Id && s.succ.Id == s.node.Id {
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
	} else if candidatePred > s.node.Id && s.succ.Id == s.node.Id {
		return &pb.Node{Id: s.node.Id, Address: s.node.Address}, nil
	} else {
		node, err := s.succ.FindPredecessor(ctx, request.Id)
		if err != nil {
			return nil, err
		}
		return &pb.Node{Id: node.Id, Address: node.Address}, nil
	}

}

// Node mutual recursion

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
