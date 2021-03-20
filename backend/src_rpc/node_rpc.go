package src_rpc

import (
	"backend/pb"
	"context"
	"errors"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	//"log"
)


type ChordNode struct{
	Id int32
	Address string
}

type ChordServer struct {
	node *ChordNode
	succ *ChordNode
	succSucc *ChordNode
	shortcuts []*ChordNode
	minSize int32
	maxSize int32
}

func (s *ChordServer) AddShortcut(ctx context.Context, node *pb.Node) (*empty.Empty, error) {
	s.shortcuts = append(s.shortcuts, &ChordNode{
		Id:      node.Id,
		Address: node.Address,
	})
	return &empty.Empty{}, nil
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
	successor, _ := s.FindSuccessor(ctx, &pb.FindSuccessorRequest{Id: s.node.Id + 1})
	s.succ = &ChordNode{Id: successor.Id, Address: successor.Address}
	// Fixing succ succ
	successorSuccessor, _ := s.FindSuccessor(ctx, &pb.FindSuccessorRequest{Id: s.succ.Id + 1})
	s.succSucc = &ChordNode{Id: successorSuccessor.Id, Address: successorSuccessor.Address}


	for idx, val := range s.shortcuts {

		conn, _ := grpc.Dial(val.Address, grpc.WithInsecure())

		c := pb.NewChordClient(conn)

		ping, _ := c.Ping(ctx, &empty.Empty{})

		if ping == nil {

			s.shortcuts = s.shortcuts[:idx+copy(s.shortcuts[idx:], s.shortcuts[idx+1:])]

		}

		_ = conn.Close()

	}

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

		//fmt.Println("Currently at: ", current.Id)

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

	if node.Id > s.maxSize || node.Id < s.minSize {

		return &empty.Empty{}, nil

	}

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

	if ShouldContainValue(s.succ.Id, request.Id, s.node.Id) {

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

	if ShouldContainValue(s.succ.Id, request.Id, s.node.Id) {

		return &pb.Node{Id: s.node.Id, Address: s.node.Address}, nil

	} else {
		node, err := s.succ.FindPredecessor(ctx, request.Id)
		if err != nil {
			return nil, err
		}
		return &pb.Node{Id: node.Id, Address: node.Address}, nil
	}


	/*
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

	 */

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

// Closest Node To

func (s *ChordServer) ClosestNodeTo(ctx context.Context, n *pb.ClosestNodeToRequest) (*pb.Node, error){

	succSuccDistance := RingDistance(s.succSucc.Id, n.Id, s.maxSize, s.minSize)

	if len(s.shortcuts) == 0 {

		return &pb.Node{
			Id:      s.succSucc.Id,
			Address: s.succSucc.Address,
		}, nil

	} else {

		smallestDistance := succSuccDistance
		closestHop := s.succSucc

		for _, shortcut := range s.shortcuts {

			shortcutDistance := RingDistance(shortcut.Id, n.Id, s.maxSize, s.minSize)

			if shortcutDistance < smallestDistance {

				smallestDistance = shortcutDistance
				closestHop = shortcut

			}

		}

		return &pb.Node{
			Id: closestHop.Id,
			Address: closestHop.Address,
		}, nil

	}

}

func (n *ChordNode) Lookup(ctx context.Context, id, hops int32) (*ChordNode, int32, error) {
	fmt.Println("Called once", hops, "at", n.Id)
	//hops = hops + 2
	fmt.Println("New hops", hops)
	conn, err := grpc.Dial(n.Address, grpc.WithInsecure())
	if err != nil {
		return nil, hops, err
	}
	defer conn.Close()
	c := pb.NewChordClient(conn)
	r, err := c.Lookup(ctx, &pb.LookupRequest{Id: id, Hops: hops})
	if err != nil {
		return nil, hops, err
	}
	return &ChordNode{r.Node.Id, r.Node.Address}, r.Hops, nil
}


func (s *ChordServer) Lookup(ctx context.Context, request *pb.LookupRequest) (*pb.LookupResponse, error) {

	if s.node.Id == request.Id {

		return &pb.LookupResponse{Node: &pb.Node{Id: s.node.Id, Address: s.node.Address}, Hops: request.Hops}, nil

	} else if request.Id > s.maxSize || request.Id < s.minSize {

		return nil, nil

	} else if ShouldContainValueTwo(s.succ.Id, request.Id, s.node.Id) {

		return &pb.LookupResponse{Node: &pb.Node{Id: s.succ.Id, Address: s.succ.Address}, Hops: request.Hops + 1}, nil

	} else if ShouldContainValueTwo(s.succSucc.Id, request.Id, s.succ.Id) {


		return &pb.LookupResponse{Node: &pb.Node{Id: s.succSucc.Id, Address: s.succSucc.Address}, Hops: request.Hops + 2}, nil

	} else {
		next, _ := s.ClosestNodeTo(ctx, &pb.ClosestNodeToRequest{Id: request.Id})


		nextNode := ChordNode{
			Id:      next.Id,
			Address: next.Address,
		}

		var nodeResponse *ChordNode
		hops := request.Hops
		var err error

		if nextNode.Id == s.succSucc.Id {



			hops = hops + 2

		} else {



			hops = hops + 1

		}
		nodeResponse, hops, err = nextNode.Lookup(ctx, request.Id, hops)


		//nodeResponse, hops, err := s.succSucc.Lookup(ctx, request.Id, request.Hops+2)
		if err != nil {
			return nil, err
		}
		return &pb.LookupResponse{Node: &pb.Node{Id: nodeResponse.Id, Address: nodeResponse.Address}, Hops: hops}, nil
	}

}

func (s *ChordServer) Leave(ctx context.Context, mpty *empty.Empty) (*empty.Empty, error) {


	if s.node.Id == s.succ.Id && s.node.Id == s.succSucc.Id {

		newError := errors.New("Oh shizzle")

		return &empty.Empty{}, newError

	}

	predecessor, _ := s.FindPredecessor(ctx, &pb.FindPredecessorRequest{Id: s.node.Id-1})
	//log.Printf("Successor: %v", successor)
	predPred, _ := s.FindPredecessor(ctx, &pb.FindPredecessorRequest{Id: predecessor.Id-1})
	// Updating Successors
	// Setting pred succ
	conn, err := grpc.Dial(predecessor.Address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Error accessing the predecessor's server")
		return &empty.Empty{}, err
	}
	c := pb.NewChordClient(conn)
	_, err = c.SetSucc(ctx, &pb.Node{Id: s.succ.Id,
		Address: s.succ.Address})
	conn.Close()
	if err != nil {
		fmt.Println("Error setting the node's successor")
		return &empty.Empty{}, err
	}

	// Setting the node's pred pred succ succ
	conn, err = grpc.Dial(predPred.Address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Error accessing the predecessor's server")
		return &empty.Empty{}, err
	}
	c = pb.NewChordClient(conn)
	_, err = c.SetSuccSucc(ctx, &pb.Node{Id: s.succ.Id,
		Address: s.succ.Address})
	conn.Close()
	if err != nil {
		fmt.Println("Error setting predecessor's successor")
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, err

}