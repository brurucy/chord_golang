package main

import (
	"backend/pb"
	"backend/src_rpc"
	"context"
	"fmt"
	prompt "github.com/c-bata/go-prompt"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"sort"
	"strconv"
	"strings"
)

// Keeping track of all currently active RPC Servers, that have some Node listening on
var grpcServers []*grpc.Server
// We are just saving the Addresses of the servers. the nodes are NOT aware of each other.
// We are using RPC in order to communicate with, and between, the nodes. It is fully multithreaded and networked.
var chordServers []*src_rpc.ChordServer



func executor(in string) {
	//fmt.Println("Your input: ")

	whitespaceSplit := strings.Fields(in)

	if whitespaceSplit[0] != "Materialize" &&
		whitespaceSplit[0] != "List" &&
		whitespaceSplit[0] != "Lookup" &&
		whitespaceSplit[0] != "Shutdown" &&
		whitespaceSplit[0] != "Join"{

		fmt.Println("Invalid command")

	} else {

		if (whitespaceSplit[0] == "Materialize" || whitespaceSplit[0] == "List" || whitespaceSplit[0] == "Shutdown") {

			if len(whitespaceSplit) > 1 {

				fmt.Println("Materialize and List take no arguments")

			} else if whitespaceSplit[0] == "Materialize"{

				chordServers, grpcServers = src_rpc.Materialize()

			} else if whitespaceSplit[0] == "List" {

				for _, val := range chordServers {

					sort.Slice(chordServers,  func(i, j int) bool {
						return chordServers[i].Node.Id < chordServers[j].Node.Id
					})

					fmt.Println(val.Node.Id, val.Shortcuts, "S-",val.Succ.Id, "NS-", val.SuccSucc.Id)

				}

			} else if whitespaceSplit[0] == "Shutdown" {

				src_rpc.ShutDown(grpcServers)

			}

		} else if whitespaceSplit[0] == "Lookup" {

			if len(whitespaceSplit) < 2 || len(whitespaceSplit) > 2 {

				fmt.Println("Lookup takes one argument, either a key or key:node")

			} else {

				keyStartNodes := strings.Split(whitespaceSplit[1], ":")
				fmt.Println(keyStartNodes)

				if len(keyStartNodes) > 2 {

					fmt.Println("There cannot be more than 2 values delimited by :")

				} else {

					if len(keyStartNodes) == 1 {

						smallest := chordServers[0]
						for _, num := range chordServers[1:] {
							if num.Node.Id < smallest.Node.Id {
								smallest = num
							}
						}

						keyId, err := strconv.ParseInt(keyStartNodes[0], 10,32)

						if int32(keyId) <= chordServers[0].Minsize || int32(keyId) > chordServers[0].Maxsize {

							fmt.Println("Node is smaller or equal than minSize or bigger than maxSize ")

						} else if err != nil {

							fmt.Println("Lookup only takes integers, optionally spaced with the : delim")

						} else {

							lookupResponse, err := smallest.Lookup(context.Background(), &pb.LookupRequest{Id: int32(keyId), Hops: 0})

							if err != nil {

								fmt.Println("Failed to Lookup")

							} else {

								fmt.Println("Data in node", lookupResponse.Node.Id, "with", lookupResponse.Hops, "hops")

							}

						}


					} else {

						keyId, err := strconv.ParseInt(keyStartNodes[0], 10,32)

						if int32(keyId) <= chordServers[0].Minsize || int32(keyId) > chordServers[0].Maxsize {

							fmt.Println("Node is smaller or equal than minSize or bigger than maxSize ")

 						} else if err == nil {

							nodeId, err := strconv.ParseInt(keyStartNodes[1], 10, 32)

							_, isNodeInTheAddressList := src_rpc.Find(chordServers, int32(nodeId))

							if isNodeInTheAddressList == false {

								fmt.Println("Can't lookup from a node that does not exist")

							} else if err == nil {

								lookupStartNode := chordServers[0]
								for _, num := range chordServers[1:] {
									if num.Node.Id == int32(nodeId) {
										lookupStartNode = num
									}
								}

								lookupResponse, err := lookupStartNode.Lookup(context.Background(), &pb.LookupRequest{Id: int32(keyId), Hops: 0})

								if err != nil {

									fmt.Println("Failed to Lookup")

								}

								fmt.Println("Data in node", lookupResponse.Node.Id, "with", lookupResponse.Hops, "hops")


							} else {

								fmt.Println("Couldn't parse the source node")

							}

						} else {

							fmt.Println("Couldn't parse the destination node", err)

						}


					}

				}

			}

		} else if whitespaceSplit[0] == "Join" {

			if len(whitespaceSplit) < 2 || len(whitespaceSplit) > 2 {

				fmt.Println("Join takes one argument, an id, not more, nor less")

			} else {

				keyId, err := strconv.ParseInt(whitespaceSplit[1], 10,32)

				fmt.Println(keyId)

				if err == nil {

					_, isNodeInTheAddressList := src_rpc.Find(chordServers, int32(keyId))

					if isNodeInTheAddressList == true {

						fmt.Println("Node is already in the ring.")

					} else if int32(keyId) <= chordServers[0].Minsize || int32(keyId) > chordServers[0].Maxsize {

						fmt.Println("Node is smaller or equal than minSize or bigger than maxSize ")

					} else {

						smallest := chordServers[0]
						for _, num := range chordServers[1:] {
							if num.Node.Id < smallest.Node.Id {
								smallest = num
							}
						}

						newNode := &src_rpc.ChordNode{
							Id:      int32(keyId),
							Address: fmt.Sprintf("127.0.0.1:%v", 10100+len(chordServers)),
						}

						// Initializing chord and gRpc server
						done := make(chan bool)
						newChordServer := &src_rpc.ChordServer{Node: newNode}
						newGrpcServer := grpc.NewServer()
						go src_rpc.RunServer(newGrpcServer, newChordServer, done)
						<-done

						chordServers = append(chordServers, newChordServer)
						grpcServers = append(grpcServers, newGrpcServer)

						_, err = smallest.Join(context.Background(), &pb.Node{Id: newNode.Id,
							Address: newNode.Address})

						if err != nil {

							fmt.Println("Failed to join the network")

						} else {

							// Recursively Propagate a message for the next node to stabilize itself, ends when next = itself
							smallest.StabilizeAll(context.Background(), &empty.Empty{})
						}

					}

				} else {

					fmt.Println("invalid input for Join")

				}

			}

		}

	}

}

func completer(in prompt.Document) []prompt.Suggest {
s := []prompt.Suggest{
{Text: "Materialize", Description: "Loads the default config, no input"},
//{Text: "Read", Description: "Reads a .txt file in the specified format"},
{Text: "List", Description: "Lists all current active nodes in the ring, no input"},
{Text: "Lookup", Description: "Lookups up a node, key:start_node"},
{Text: "Join", Description: "Joins the given node Id with the ring"},
{Text: "Shutdown", Description: "Shuts down the whole cluster"},
}
return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func main() {
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix("λ "),
		prompt.OptionTitle("prompt for huber's take on chord"),
	)
	p.Run()
}