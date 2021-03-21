package main

import (
	"chord_golang/parser"
	"chord_golang/pb"
	"chord_golang/src_rpc"
	"context"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"

	prompt "github.com/c-bata/go-prompt"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

// Keeping track of all currently active RPC Servers, that have some Node listening on
var grpcServers []*grpc.Server

// We are just saving the Addresses of the servers. the nodes are NOT aware of each other.
// We are using RPC in order to communicate with, and between, the nodes. It is fully multithreaded and networked.
var chordServers []*src_rpc.ChordServer
var MinSize int32
var MaxSize int32

func executor(in string) {

	whitespaceSplit := strings.Fields(in)

	if len(whitespaceSplit) == 0 {
		return
	}

	if whitespaceSplit[0] != "Read" &&
		whitespaceSplit[0] != "Materialize" &&
		whitespaceSplit[0] != "List" &&
		whitespaceSplit[0] != "Lookup" &&
		whitespaceSplit[0] != "Shutdown" &&
		whitespaceSplit[0] != "Join" &&
		whitespaceSplit[0] != "Leave" &&
		whitespaceSplit[0] != "Shortcut" {

		fmt.Println("Invalid command")

	} else {

		if whitespaceSplit[0] == "Materialize" || whitespaceSplit[0] == "List" || whitespaceSplit[0] == "Shutdown" {

			if len(whitespaceSplit) > 1 {

				fmt.Println("Materialize and List take no arguments")

			} else if whitespaceSplit[0] == "Materialize" {

				chordServers, grpcServers = src_rpc.Materialize()

				MinSize = chordServers[0].Minsize
				MaxSize = chordServers[0].Maxsize

			} else if whitespaceSplit[0] == "List" {

				copyChordServers := append([]*src_rpc.ChordServer{}, chordServers...)

				sort.Slice(copyChordServers, func(i, j int) bool {
					return copyChordServers[i].Node.Id < copyChordServers[j].Node.Id
				})

				for _, val := range copyChordServers {

					var shortcuts []int32

					for _, val := range val.Shortcuts {

						shortcuts = append(shortcuts, val.Id)

					}

					fmt.Println(val.Node.Id, shortcuts, "S-", val.Succ.Id, "NS-", val.SuccSucc.Id, "on port:", val.Node.Address)
					//fmt.Println(val.Node.Id)
					//fmt.Println(shortcuts)
					//fmt.Println("S-", val.Succ.Id)
					//fmt.Println("NS-", val.SuccSucc.Id)
					//fmt.Println("on port:", val.Node.Address)

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

						keyId, err := strconv.ParseInt(keyStartNodes[0], 10, 32)

						if int32(keyId) <= MinSize || int32(keyId) > MaxSize {

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

						keyId, err := strconv.ParseInt(keyStartNodes[0], 10, 32)

						if int32(keyId) <= MinSize || int32(keyId) > MaxSize {

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

				keyId, err := strconv.ParseInt(whitespaceSplit[1], 10, 32)

				fmt.Println(keyId)

				if err == nil {

					_, isNodeInTheAddressList := src_rpc.Find(chordServers, int32(keyId))

					if isNodeInTheAddressList == true {

						fmt.Println("Node is already in the ring.")

					} else if int32(keyId) <= MinSize || int32(keyId) > MaxSize {

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
							Address: fmt.Sprintf("127.0.0.1:%v", 10000+rand.Intn(10000)),
						}

						// Initializing chord and gRpc server
						done := make(chan bool)
						newChordServer := &src_rpc.ChordServer{Node: newNode, Minsize: MinSize, Maxsize: MaxSize}
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
							fmt.Println("Stabilize attempt")
							smallest.StabilizeAll(context.Background(), &empty.Empty{})
							fmt.Println("Stabilize completed")
						}

					}

				} else {

					fmt.Println("invalid input for Join")

				}

			}

		} else if whitespaceSplit[0] == "Leave" {

			if len(whitespaceSplit) < 2 || len(whitespaceSplit) > 2 {

				fmt.Println("Leave takes one argument, an id, not more, nor less")

			} else {

				keyId, err := strconv.ParseInt(whitespaceSplit[1], 10, 32)

				if err == nil {

					if int32(keyId) <= MinSize || int32(keyId) > MaxSize || len(chordServers) == 1 {

						fmt.Println("Node is smaller or equal than minSize or bigger than maxSize or trying to leave as the last in the ring")

					} else {

						pos, isNodeInTheAddressList := src_rpc.Find(chordServers, int32(keyId))

						if isNodeInTheAddressList == false {

							fmt.Println("Can't Leave from a node that does not exist")

						} else {

							leaveServer := &(chordServers[pos])
							leaveGrpcServer := &(grpcServers[pos])

							var stabilizerNodeId int32

							/*
								for i, _ := range chordServers {

									if i != pos {

										stabilizerNodeId = int32(i)

									}

								}
							*/

							stabilizerNodeId = chordServers[pos].Succ.Id

							stabilizerNodeIndex, _ := src_rpc.Find(chordServers, stabilizerNodeId)

							fmt.Println("Attempting to leave at", keyId, "with stabilizer", stabilizerNodeId) //chordServers[stabilizerNodeId].Node.Id)
							(*leaveServer).Leave(context.Background(), &empty.Empty{})
							fmt.Println("Left")
							(*leaveGrpcServer).Stop()
							fmt.Println("Successfully stopped gRPC server")

							if len(chordServers) < 2 {

								(*chordServers[pos]).Succ = (*chordServers[pos]).Node
								(*chordServers[pos]).SuccSucc = (*chordServers[pos]).Node

							} else {

								chordServers[stabilizerNodeIndex].StabilizeAll(context.Background(), &empty.Empty{})
								fmt.Println("Succesfully stabilized, on nodeId ", stabilizerNodeId)

							}

							for idx, val := range chordServers {

								conn, _ := grpc.Dial(val.Node.Address, grpc.WithInsecure())

								c := pb.NewChordClient(conn)

								ping, _ := c.Ping(context.Background(), &empty.Empty{})

								//fmt.Println("Pinging", val.Node.Id, "at",val.Node.Address)

								_ = conn.Close()

								if ping == nil {

									chordServers[idx] = nil
									grpcServers[idx] = nil
									fmt.Println(val.Node.Id, "is not alive")

								}

							}

							n := 0
							for idx, x := range chordServers {
								if x != nil {
									chordServers[n] = x
									grpcServers[n] = grpcServers[idx]
									n++
								}
							}
							chordServers = chordServers[:n]
							grpcServers = grpcServers[:n]

						}
					}
				} else {

					fmt.Println("Failed to parse the which node to leave")

				}
			}
		} else if whitespaceSplit[0] == "Shortcut" {

			if len(whitespaceSplit) < 2 || len(whitespaceSplit) > 2 {

				fmt.Println("Shortcut takes one argument, an id, not more, nor less")

			} else {

				keyStartNodes := strings.Split(whitespaceSplit[1], ":")
				fmt.Println(keyStartNodes)

				if len(keyStartNodes) != 2 {

					fmt.Println("There cannot be more, or less, than 2 values delimited by :")

				} else {

					keyId, err := strconv.ParseInt(keyStartNodes[0], 10, 32)

					if err != nil {

						fmt.Println("Failed to parse the node")

					} else {

						shortcutId, err := strconv.ParseInt(keyStartNodes[1], 10, 32)

						if err != nil {

							fmt.Println("Failed to parse the shortcut")

						} else {

							pos_node, isNodeInTheAddressList := src_rpc.Find(chordServers, int32(keyId))

							if isNodeInTheAddressList == false {

								fmt.Println("Can't add a shortcut to a node that does not exist")

							} else {

								pos_shortcut, isShortcutInTheAddressList := src_rpc.Find(chordServers, int32(shortcutId))

								if isShortcutInTheAddressList == false {

									fmt.Println("Can't add a shortcut that does not exist")

								} else {

									_, err = chordServers[pos_node].AddShortcut(context.Background(), &pb.Node{Id: chordServers[pos_shortcut].Node.Id, Address: chordServers[pos_shortcut].Node.Address})

									if err != nil {

										fmt.Println("Error adding to list of shortcuts")

									}

								}

							}

						}

					}

				}

			}

		} else if whitespaceSplit[0] == "Read" {

			if len(whitespaceSplit) < 2 || len(whitespaceSplit) > 2 {

				fmt.Println("Read takes one argument, a text file")

			} else {

				file_location := whitespaceSplit[1]

				file, err := parser.Parse(file_location)

				if err != nil {

					fmt.Printf("Something bad happened whilst parsing", err)

				} else {

					MinSize = file.MinSize
					MaxSize = file.MaxSize
					nodes := file.Nodes
					shortcuts := file.Shortcuts

					chordServers, grpcServers = src_rpc.Read(MinSize, MaxSize, nodes, shortcuts)

				}

			}

		}
	}
}

func completer(in prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		//{Text: "Materialize", Description: "Loads the default config, no input needed"},
		//{Text: "Read", Description: "Reads a .txt file in the specified format"},
		{Text: "List", Description: "Lists all current active nodes in the ring, no input"},
		{Text: "Lookup", Description: "Lookups up a node, key:start_node"},
		{Text: "Join", Description: "Joins the given node Id with the ring"},
		{Text: "Leave", Description: "Shuts down the specified Node"},
		{Text: "Shortcut", Description: "Adds a shortcut to the specified node"},
		//{Text: "Shutdown", Description: "Shuts down the whole cluster"},
	}
	return prompt.FilterHasPrefix(s, in.GetWordBeforeCursor(), true)
}

func main() {

	if len(os.Args) < 2 {

		fmt.Println("No argument given")
		os.Exit(1)

	} else {

		parsedArgs := strings.TrimSpace(os.Args[1])

		if strings.HasSuffix(parsedArgs, ".txt") {

			executor(fmt.Sprintf("Read %s", parsedArgs))

			p := prompt.New(
				executor,
				completer,
				prompt.OptionPrefix("λ "),
				prompt.OptionTitle("prompt for huber's take on chord"),
			)
			p.Run()
		} else {

			fmt.Println("Please provide a .txt file")
			os.Exit(1)

		}
	}
}
