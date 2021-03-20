package main

import (
	"backend/pb"
	"backend/src_rpc"
	"context"
	"fmt"
	prompt "github.com/c-bata/go-prompt"
	"google.golang.org/grpc"
	"strconv"
	"strings"
)

var grpcServers []*grpc.Server
// We are just saving the Addresses of the servers. the nodes are NOT aware of each other.
// We are using RPC in order to communicate with, and between, the nodes.
var chordServers []*src_rpc.ChordServer



func executor(in string) {
	//fmt.Println("Your input: ")

	whitespaceSplit := strings.Fields(in)

	if whitespaceSplit[0] != "Materialize" && whitespaceSplit[0] != "List" && whitespaceSplit[0] != "Lookup" && whitespaceSplit[0] != "Shutdown" {

		fmt.Println("Invalid command")

	} else {

		if (whitespaceSplit[0] == "Materialize" || whitespaceSplit[0] == "List" || whitespaceSplit[0] == "Shutdown") {

			if len(whitespaceSplit) > 1 {

				fmt.Println("Materialize and List take no arguments")

			} else if whitespaceSplit[0] == "Materialize"{

				chordServers, grpcServers = src_rpc.Materialize()

			} else if whitespaceSplit[0] == "List" {

				for _, val := range chordServers {

					fmt.Println(val.Node.Id, val.Succ.Id, val.SuccSucc.Id)

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

						smallestId, err := strconv.ParseInt(keyStartNodes[0], 10,32)

						if err != nil {

							fmt.Println("Lookup only takes integers, optionally spaced with the : delim")

						} else {

							lookupResponse, err := smallest.Lookup(context.Background(), &pb.LookupRequest{Id: int32(smallestId), Hops: 0})

							if err != nil {

								fmt.Println("Failed to Lookup")

							} else {

								fmt.Println("Data in node", lookupResponse.Node.Id, "with", lookupResponse.Hops, "hops")

							}

						}


					} else {

						keyId, err := strconv.ParseInt(keyStartNodes[0], 10,32)

						if err == nil {

							nodeId, err := strconv.ParseInt(keyStartNodes[1], 10, 32)

							_, isNodeInTheAddressList := src_rpc.Find(chordServers, int32(nodeId))

							if isNodeInTheAddressList == false {

								fmt.Println("Can't lookup from a node that does not exist")

							} else if err == nil {

								lookupStartNode := chordServers[0]            // set the smallest number to the first element of the list
								for _, num := range chordServers[1:] { // iterate over the rest of the list
									if num.Node.Id == int32(nodeId) {     // if num is smaller than the current smallest number
										lookupStartNode = num      // set smallest to num
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

		}

	}

}

func completer(in prompt.Document) []prompt.Suggest {
s := []prompt.Suggest{
{Text: "Materialize", Description: "Loads the default config, no input"},
//{Text: "Read", Description: "Reads a .txt file in the specified format"},
{Text: "List", Description: "Lists all current active nodes, no input"},
{Text: "Lookup", Description: "Lookups up a node, key:start_node"},
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