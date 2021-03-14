package src

import (
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
	Ring              *Ring
}

func AbsInt(x int) int {

	if x < 0 {

		return -x

	}
	return x

}

func Mod(a, b int) int {
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
	maxSizeToFrom := AbsInt((maxSize - (to - from)))

	result := Min(toFrom, maxSizeToFrom)

	if to > from {

		return result

	} else if from == to {

		return 0

	} else {

		result = Mod((maxSize - minSize + result), maxSize)
		return result

	}

}

// Find returns the smallest index i at which x == a[i],
// or len(a) if there is no such index.
func Find(a []string, x string) int {
	for i, n := range a {
		if x == n {
			return i
		}
	}
	return len(a)
}

func (n *Node) NextClosestHopTo(key int) *Node {

	// we don't need to consider Succ because we know that SuccSucc doesn't have the value
	// TODO: what if ring destabilizes and both Suc and SuccSucc don't have value
	// but actually the key is between Succ and SuccSucc

	// fmt.Println("Calling RingDistance with args", n.SuccSucc.Id, key, n.Ring.MaxSize, n.Ring.MinSize)
	succSuccDistance := RingDistance(n.SuccSucc.Id, key, n.Ring.MaxSize, n.Ring.MinSize)
	// fmt.Println("Succsucc", n.SuccSucc.Id, "has distance:", succSuccDistance, "to key:", key)

	if len(n.Shortcuts) == 0 {

		// fmt.Println("No shortcuts found, just giving you SuccSucc")

		return n.SuccSucc

	} else {

		// remember best hop among direct successors
		smallestDistance := succSuccDistance
		closestHop := n.SuccSucc

		// fmt.Println("Shortcuts found.")

		// check if shortcuts give a better hop than successors
		for _, shortcut := range n.Shortcuts {
			shortcutDistance := RingDistance(shortcut.Id, key, n.Ring.MaxSize, n.Ring.MinSize)
			// fmt.Println("Considering a shortcut through:", shortcut.Id, "with distance", shortcutDistance)

			if shortcutDistance < smallestDistance {
				// fmt.Println("Better path found through shortcut:", shortcut.Id, "with distance", shortcutDistance)
				smallestDistance = shortcutDistance
				closestHop = shortcut
			}
		}

		return closestHop
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

// FIXME: It most probably doesn't work. Will deal with it later
func (n *Node) FindPredecessor(key int) *Node {

	/*
		fmt.Println("Currently at: ", n.Id)
		fmt.Println("Next: ", n.Succ.Id)
		fmt.Println("Next Next: ", n.SuccSucc.Id)
	*/

	if n.Id < key && key <= n.Succ.Id {

		return n

	} else if n.Id < key && n.Id > n.Succ.Id && key <= n.Succ.Id { // key=5 n=4 suc=1

		return n

	} else if n.Id < key && n.Id > n.Succ.Id && key >= n.Succ.Id {

		return n

	} else if n.Id > key && n.Id > n.Succ.Id && key <= n.Succ.Id {

		return n

	} else {

		nextHop := n.NextClosestHopTo(key)
		return nextHop.Lookup(key)

	}

}

func (n *Node) HasValue(key int) bool {
	return n.Values[key]
}

func ShouldContainValue(id int, key int, predId int) bool {
	return id >= key && key > predId
}

func (n *Node) Lookup(key int) *Node {
	var emptyNode Node

	// See if we have value already
	if n.Values[key] {
		return n
	}

	// should succ have value?
	if ShouldContainValue(n.Succ.Id, key, n.Id) {

		// then ask her!
		if n.Succ.HasValue(key) {
			return n.Succ
		} else {
			return &emptyNode // TODO: define behavior when key not found
		}
	}

	// should succsucc have value?
	if ShouldContainValue(n.SuccSucc.Id, key, n.Succ.Id) {

		// then ask her!
		if n.SuccSucc.HasValue(key) {
			return n.SuccSucc
		} else {
			return &emptyNode // TODO: define behavior when key not found
		}
	}

	nextHop := n.NextClosestHopTo(key)
	return nextHop.Lookup(key)
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

	pred := (*(n.FindPredecessor(n.Id))).Id

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
