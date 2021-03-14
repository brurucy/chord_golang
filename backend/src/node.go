package src

import (
	"fmt"
	"sort"
)

type Node struct {
	Id                int
	Succ              *Node
	SuccSucc          *Node
	Successors        []*Node
	Shortcuts         []*Node
	StabilizationRate int
	Values            map[int]bool
}

func AbsInt(x int) int {

	if x < 0 {

		return -x

	}
	return x

}

func mod(a, b int) int {
	m := a % b
	if a < 0 && b < 0 {
		m -= b
	}
	if a < 0 && b > 0 {
		m += b
	}
	return m
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func RingDistance(from, to, maxSize, minSize int) int {

	toFrom := to - from
	maxSizeToFrom := AbsInt((maxSize - to - from))

	result := Min(toFrom, maxSizeToFrom)

	if to > from {

		return result

	} else if from == to {

		return 0

	} else {

		result = mod((maxSize - minSize + result), maxSize)
		return result

	}

}

func (n *Node) ClosestHopTo(key int) *Node {

	if len(n.Shortcuts) == 0 {

		succDistance := AbsInt((n.Succ).Id - key)
		succSuccDistance := AbsInt((n.SuccSucc).Id - key)

		if key > n.Id {

			if succDistance > succSuccDistance && n.Succ.Id < n.Id {

				return n.Succ

			}

			return n.SuccSucc

		} else {

			if succDistance > succSuccDistance {

				fmt.Println(succDistance, succSuccDistance)

				return n.Succ

			}

			return n.SuccSucc

		}

	} else {

		var possiblyClosestNode *Node

		// This is doing a binary search over shortcuts
		closestValue := sort.Search(len(n.Shortcuts), func(i int) bool { return n.Shortcuts[i].Id >= key })

		if closestValue == len(n.Shortcuts) {

			possiblyClosestNode = n.Shortcuts[closestValue-1]

		} else {

			possiblyClosestNode = n.Shortcuts[closestValue]

		}

		succDistance := AbsInt(n.Succ.Id - key)
		succSuccDistance := AbsInt(n.SuccSucc.Id - key)
		possiblyClosestNodeDistance := AbsInt(possiblyClosestNode.Id - key)

		distSlice := make(map[int]*Node, 3)
		distSlice[succDistance] = n.Succ
		distSlice[succSuccDistance] = n.SuccSucc
		distSlice[possiblyClosestNodeDistance] = possiblyClosestNode

		closest := n.Succ
		closestIdx := succDistance

		if key > n.Id {

			for idx, val := range distSlice {

				if idx < closestIdx {

					closest = val
					closestIdx = idx

				}

			}

		} else {

			for idx, val := range distSlice {

				if idx > closestIdx {

					closest = val
					closestIdx = idx

				}

			}

		}

		return closest

	}

}

func (n *Node) AddShortcut(shortcut *Node) {

	// Just leave it linear, no need for bisection
	//n.Shortcuts = append(n.Shortcuts, shortcut)

	candidateIndex := sort.Search(len(n.Shortcuts), func(i int) bool { return n.Shortcuts[i].Id >= (*shortcut).Id })

	// Insert, same for every insert
	n.Shortcuts = append(n.Shortcuts, &Node{})
	copy(n.Shortcuts[candidateIndex+1:], n.Shortcuts[candidateIndex:])
	n.Shortcuts[candidateIndex] = shortcut

}

func (n *Node) findPredecessor(key int) *Node {

	/*
		fmt.Println("Currently at: ", n.Id)
		fmt.Println("Next: ", n.Succ.Id)
		fmt.Println("Next Next: ", n.SuccSucc.Id)
	*/

	if n.Id < key && key <= n.Succ.Id {

		return n

	} else if n.Id < key && n.Id > n.Succ.Id && key <= n.Succ.Id {

		return n

	} else if n.Id < key && n.Id > n.Succ.Id && key >= n.Succ.Id {

		return n

	} else if n.Id > key && n.Id > n.Succ.Id && key <= n.Succ.Id {

		return n

	} else {

		nextHop := n.ClosestHopTo(key)
		return nextHop.findValue(key)

	}

}

func (n *Node) findValue(key int) *Node {

	// Just check if the key is in current node's VALUE set

	if n.Id < key && key <= n.Succ.Id {

		return n.Succ

	} else if n.Id < key && n.Id > n.Succ.Id && key <= n.Succ.Id {

		return n.Succ

	} else if n.Id < key && n.Id > n.Succ.Id && key >= n.Succ.Id {

		return n.Succ

	} else if n.Id > key && n.Id > n.Succ.Id && key <= n.Succ.Id {

		return n.Succ

	} else {

		nextHop := n.ClosestHopTo(key)
		return nextHop.findValue(key)

	}

}

func (n *Node) Stabilize() {

	// Successors

	currentSucc := n.Succ
	currentSuccSucc := n.SuccSucc

	for {

		if currentSucc == nil {

			n.Succ = n.SuccSucc
			n.SuccSucc = (n.SuccSucc).Succ

			currentSucc = n.Succ
			currentSuccSucc = n.SuccSucc

		} else if currentSuccSucc == nil {

			n.SuccSucc = (*n.Succ).Succ
			currentSuccSucc = n.SuccSucc

		}

		if currentSucc != nil && currentSuccSucc != nil {

			break

		}

	}

	// Shortcuts

	counter := 0
	for _, val := range n.Shortcuts {

		if val != nil {

			n.Shortcuts[counter] = val
			counter++

		}
	}
	n.Shortcuts = n.Shortcuts[:counter]

}

func (n *Node) MigrateData(maxRange int) {

	values := make(map[int]bool)

	pred := (*(n.findPredecessor(n.Id))).Id

	if pred < n.Id {

		for i := pred + 1; i <= n.Id; i++ {

			values[i] = true

		}

	} else {

		// First half

		for i := pred + 1; i <= maxRange; i++ {

			values[i] = true

		}

		for i := 0; i <= n.Id; i++ {

			values[i] = true

		}

	}

	n.Values = values

}
