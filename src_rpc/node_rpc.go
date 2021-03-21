package src_rpc

import (
	"chord_golang/pb"
	"context"
	"errors"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	//"log"
)

type ChordNode struct {
	Id      int32
	Address string
}

type ChordServer struct {
	Node      *ChordNode
	Succ      *ChordNode
	SuccSucc  *ChordNode
	Shortcuts []*ChordNode
	Data      map[int32]bool
	Minsize   int32
	Maxsize   int32
}

func (s *ChordServer) AddShortcut(ctx context.Context, node *pb.Node) (*empty.Empty, error) {
	s.Shortcuts = append(s.Shortcuts, &ChordNode{
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

func (s *ChordServer) MigrateData(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {

	currentId := s.Node.Id
	pred, _ := s.FindPredecessor(ctx, &pb.FindPredecessorRequest{Id: s.Node.Id})
	s.Data = make(map[int32]bool, 0)
	predId := pred.Id

	if currentId > predId {

		for i := predId + 1; i <= currentId; i++ {

			s.Data[i] = true

		}

	} else {

		for i := predId + 1; i <= s.Maxsize; i++ {

			s.Data[i] = true

		}
		for i := s.Minsize + 1; i <= currentId; i++ {

			s.Data[i] = true

		}

	}

	return e, nil

}

func (s *ChordServer) GetSucc(context.Context, *empty.Empty) (*pb.Node, error) {

	return &pb.Node{Id: s.Succ.Id, Address: s.Succ.Address}, nil

}

func (s *ChordServer) MigrateDataAll(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {

	startId := s.Node.Id
	_, err := s.MigrateData(ctx, e)
	if err != nil {
		fmt.Printf("Failed to Migrate starting node data: %v", err)
		return e, err
	}
	//fmt.Println("Passed the stabilize call")
	current := s.Succ

	for current.Id != startId {

		//fmt.Println("Currently at: ", current.Id)

		conn, err := grpc.Dial(current.Address, grpc.WithInsecure())
		if err != nil {
			fmt.Printf("Failed to Dial %v", err)
			return &empty.Empty{}, nil
		}
		c := pb.NewChordClient(conn)
		_, err = c.MigrateData(ctx, &empty.Empty{})
		if err != nil {
			fmt.Printf("Failed to Migrate %v", err)
			return &empty.Empty{}, nil
		}

		//TODO GetSucc rpc
		//next, _ := c.FindSuccessor(ctx, &pb.FindSuccessorRequest{Id: current.Id + 1})
		next, _ := c.GetSucc(ctx, &empty.Empty{})

		_ = conn.Close()

		current = &ChordNode{Id: next.Id, Address: next.Address}
	}

	return &empty.Empty{}, nil

}

func (s *ChordServer) SetSucc(ctx context.Context, node *pb.Node) (*empty.Empty, error) {
	s.Succ = &ChordNode{
		Id:      node.Id,
		Address: node.Address,
	}
	return &empty.Empty{}, nil
}

func (s *ChordServer) SetSuccSucc(ctx context.Context, node *pb.Node) (*empty.Empty, error) {
	s.SuccSucc = &ChordNode{
		Id:      node.Id,
		Address: node.Address,
	}
	return &empty.Empty{}, nil
}

func (s *ChordServer) Stabilize(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {

	// Fixing succ
	successor, _ := s.FindSuccessor(ctx, &pb.FindSuccessorRequest{Id: s.Node.Id + 1})
	s.Succ = &ChordNode{Id: successor.Id, Address: successor.Address}
	// Fixing succ succ
	successorSuccessor, _ := s.FindSuccessor(ctx, &pb.FindSuccessorRequest{Id: s.Succ.Id + 1})
	s.SuccSucc = &ChordNode{Id: successorSuccessor.Id, Address: successorSuccessor.Address}

	for idx, val := range s.Shortcuts {

		conn, _ := grpc.Dial(val.Address, grpc.WithInsecure())

		c := pb.NewChordClient(conn)

		ping, _ := c.Ping(ctx, &empty.Empty{})

		if ping == nil {

			s.Shortcuts = s.Shortcuts[:idx+copy(s.Shortcuts[idx:], s.Shortcuts[idx+1:])]

		}

		_ = conn.Close()

	}

	return &empty.Empty{}, nil
}

func (s *ChordServer) StabilizeAll(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {

	startId := s.Node.Id
	_, err := s.Stabilize(ctx, e)
	if err != nil {
		fmt.Printf("Failed to Stabilize starting node: %v", err)
		return &empty.Empty{}, err
	}
	//fmt.Println("Passed the stabilize call")
	current := s.Succ

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

	_, err = s.MigrateDataAll(ctx, &empty.Empty{})

	return &empty.Empty{}, nil

}

func (s *ChordServer) Join(ctx context.Context, node *pb.Node) (*empty.Empty, error) {

	if node.Id > s.Maxsize || node.Id < s.Minsize {

		return &empty.Empty{}, nil

	}

	successor, _ := s.FindSuccessor(ctx, &pb.FindSuccessorRequest{Id: node.Id})

	predecessor, _ := s.FindPredecessor(ctx, &pb.FindPredecessorRequest{Id: node.Id})

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

	if ShouldContainValue(s.Succ.Id, request.Id, s.Node.Id) {

		return &pb.Node{Id: s.Succ.Id, Address: s.Succ.Address}, nil

	} else if s.Node.Id == s.Succ.Id && s.Node.Id == s.SuccSucc.Id {

		return &pb.Node{Id: s.Succ.Id, Address: s.Succ.Address}, nil

	} else {
		node, err := s.Succ.FindSuccessor(ctx, request.Id)
		if err != nil {
			return nil, err
		}
		return &pb.Node{Id: node.Id, Address: node.Address}, nil
	}

}

func (s *ChordServer) FindPredecessor(ctx context.Context, request *pb.FindPredecessorRequest) (*pb.Node, error) {

	if ShouldContainValue(s.Succ.Id, request.Id, s.Node.Id) {

		return &pb.Node{Id: s.Node.Id, Address: s.Node.Address}, nil

	} else if s.Node.Id == s.Succ.Id && s.Node.Id == s.SuccSucc.Id {

		return &pb.Node{Id: s.Node.Id, Address: s.Node.Address}, nil

	} else {
		node, err := s.Succ.FindPredecessor(ctx, request.Id)
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

// Closest Node To

func (s *ChordServer) ClosestNodeTo(ctx context.Context, n *pb.ClosestNodeToRequest) (*pb.Node, error) {

	succSuccDistance := RingDistance(s.SuccSucc.Id, n.Id, s.Maxsize, s.Minsize)

	if len(s.Shortcuts) == 0 {

		return &pb.Node{
			Id:      s.SuccSucc.Id,
			Address: s.SuccSucc.Address,
		}, nil

	} else {

		smallestDistance := succSuccDistance
		closestHop := s.SuccSucc

		for _, shortcut := range s.Shortcuts {

			shortcutDistance := RingDistance(shortcut.Id, n.Id, s.Maxsize, s.Minsize)

			if shortcutDistance < smallestDistance {

				smallestDistance = shortcutDistance
				closestHop = shortcut

			}

		}

		return &pb.Node{
			Id:      closestHop.Id,
			Address: closestHop.Address,
		}, nil

	}

}

func (n *ChordNode) Lookup(ctx context.Context, id, hops int32) (*ChordNode, int32, error) {
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

func (s *ChordServer) HasValue(ctx context.Context, id int32) bool {

	/*
		pred, _ := s.FindPredecessor(ctx, &pb.FindPredecessorRequest{
			Id: s.Node.Id,
		})

		return ShouldContainValue(s.Node.Id, id, pred.Id)
	*/

	return s.Data[id]

}

func (s *ChordServer) Lookup(ctx context.Context, request *pb.LookupRequest) (*pb.LookupResponse, error) {

	next, _ := s.ClosestNodeTo(ctx, &pb.ClosestNodeToRequest{Id: request.Id})

	nextNode := ChordNode{
		Id:      next.Id,
		Address: next.Address,
	}

	if s.Node.Id == request.Id || s.HasValue(ctx, request.Id) {

		return &pb.LookupResponse{Node: &pb.Node{Id: s.Node.Id, Address: s.Node.Address}, Hops: request.Hops}, nil

	} else if request.Id > s.Maxsize || request.Id < s.Minsize {

		return nil, nil

	} else if ShouldContainValueTwo(s.Succ.Id, request.Id, s.Node.Id) {

		return &pb.LookupResponse{Node: &pb.Node{Id: s.Succ.Id, Address: s.Succ.Address}, Hops: request.Hops + 1}, nil

	} else if ShouldContainValueTwo(s.SuccSucc.Id, request.Id, s.Succ.Id) && s.SuccSucc.Id != nextNode.Id {

		return &pb.LookupResponse{Node: &pb.Node{Id: s.SuccSucc.Id, Address: s.SuccSucc.Address}, Hops: request.Hops + 2}, nil

	} else {

		/*
			next, _ := s.ClosestNodeTo(ctx, &pb.ClosestNodeToRequest{Id: request.Id})

			nextNode := ChordNode{
				Id:      next.Id,
				Address: next.Address,
			}
		*/

		var nodeResponse *ChordNode
		hops := request.Hops
		var err error

		_, isSuccSuccInsideShortcuts := FindShortcut(s.Shortcuts, s.SuccSucc.Id)

		if nextNode.Id == s.SuccSucc.Id && isSuccSuccInsideShortcuts != true {

			hops = hops + 2

		} else {

			hops = hops + 1

		}
		nodeResponse, hops, err = nextNode.Lookup(ctx, request.Id, hops)

		//nodeResponse, hops, err := s.SuccSucc.Lookup(ctx, request.Id, request.Hops+2)
		if err != nil {
			return nil, err
		}
		return &pb.LookupResponse{Node: &pb.Node{Id: nodeResponse.Id, Address: nodeResponse.Address}, Hops: hops}, nil
	}

}

func (s *ChordServer) Leave(ctx context.Context, mpty *empty.Empty) (*empty.Empty, error) {

	if s.Node.Id == s.Succ.Id && s.Node.Id == s.SuccSucc.Id {

		newError := errors.New("Oh shizzle")

		return &empty.Empty{}, newError

	}

	predecessor, _ := s.FindPredecessor(ctx, &pb.FindPredecessorRequest{Id: s.Node.Id})
	//log.Printf("Successor: %v", successor)
	predPred, _ := s.FindPredecessor(ctx, &pb.FindPredecessorRequest{Id: predecessor.Id})
	// Updating Successors
	// Setting pred succ
	conn, err := grpc.Dial(predecessor.Address, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Error accessing the predecessor's server")
		return &empty.Empty{}, err
	}
	c := pb.NewChordClient(conn)
	_, err = c.SetSucc(ctx, &pb.Node{Id: s.Succ.Id,
		Address: s.Succ.Address})
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
	_, err = c.SetSuccSucc(ctx, &pb.Node{Id: s.Succ.Id,
		Address: s.Succ.Address})
	conn.Close()
	if err != nil {
		fmt.Println("Error setting predecessor's successor")
		return &empty.Empty{}, err
	}

	return &empty.Empty{}, err

}
