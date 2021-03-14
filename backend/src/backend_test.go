package src

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRingDistance(t *testing.T) {
	assert := assert.New(t)

	from := []int{22, 4, 4, 22, 4, 4}
	to := []int{4, 22, 4, 4, 22, 4}
	maxSize := []int{100, 100, 100, 100, 100, 100}
	minSize := []int{0, 0, 0, 1, 1, 1}
	correctResults := []int{82, 18, 0, 81, 18, 0}

	var distance int

	for i := 0; i < 6; i++ {
		distance = RingDistance(from[i], to[i], maxSize[i], minSize[i])
		debugStr := fmt.Sprintf("from: %d to: %d dist: %d != %d", from[i], to[i], distance, correctResults[i])
		assert.Equal(distance, correctResults[i], debugStr)
	}

}

func GetNodeIndexById(id int, nodes []Node) int {
	for k, v := range nodes {
		if id == v.Id {
			return k
		}
	}
	return -1
}

func TestStabilizeAddShortcut(t *testing.T) {
	assert := assert.New(t)

	const minrange = 1
	const maxrange = 100
	nodeIds := []int{5, 17, 22, 56, 71, 89, 92} // 110

	ring := Ring{MinSize: minrange, MaxSize: maxrange}
	var nodes []Node

	// create nodes
	for i := 0; i < len(nodeIds); i++ {
		nodes = append(nodes, Node{Id: nodeIds[i], Ring: &ring})
	}

	// add nodes' successors
	for i := 0; i < len(nodeIds); i++ {
		next_in_ring_i := (i + 1) % len(nodeIds)
		nodes[i].Succ = &nodes[next_in_ring_i]
	}

	// stabilize nodes
	for i := 0; i < len(nodeIds); i++ {
		nodes[i].Stabilize()
	}

	// add some shortcuts
	nodes[GetNodeIndexById(5, nodes)].AddShortcut(&nodes[GetNodeIndexById(56, nodes)])
	nodes[GetNodeIndexById(5, nodes)].AddShortcut(&nodes[GetNodeIndexById(71, nodes)])
	nodes[GetNodeIndexById(22, nodes)].AddShortcut(&nodes[GetNodeIndexById(89, nodes)])
	nodes[GetNodeIndexById(92, nodes)].AddShortcut(&nodes[GetNodeIndexById(56, nodes)])

	// FIXME: stopped working after rewriting `NextClosestHop` and `Lookup` methods
	// TODO: run for all nodes?
	// migrate data
	// for i := 0; i < 2; /*len(nodeIds)*/ i++ {
	// 	nodes[i].MigrateData(maxrange)
	// }

	// test shortcuts adding
	assert.Equal(len(nodes[0].Shortcuts), 2, "should add shortcuts")

	// test stabilization
	assert.Equal(nodes[len(nodes)-1].Succ, &nodes[0], "last node in ring should have first node as Succ")
	assert.Equal(nodes[0].Succ.Succ, &nodes[2], "should find and add SuccSucc after stabilize")

	// test ClosestHops through SuccSucc
	assert.Equal(89, nodes[GetNodeIndexById(22, nodes)].NextClosestHopTo(56).Id, "should find closest hop. [22 hops 56 should find 89]")
	assert.Equal(89, nodes[GetNodeIndexById(22, nodes)].NextClosestHopTo(50).Id, "should find closest hop. [22 hops 50 should find 89]")

	assert.Equal(56, nodes[GetNodeIndexById(92, nodes)].NextClosestHopTo(4).Id, "should find closest hop. [92 hops 4 should find 17]")
	assert.Equal(17, nodes[GetNodeIndexById(92, nodes)].NextClosestHopTo(93).Id, "should find closest hop. [92 hops 93 should find 17]")

	// test ClosestHops through SuccSucc
	assert.Equal(56, nodes[GetNodeIndexById(92, nodes)].NextClosestHopTo(10).Id, "should find closest hop. [92 hops 10 should find 17]")
	assert.Equal(17, nodes[GetNodeIndexById(92, nodes)].NextClosestHopTo(19).Id, "should find closest hop. [92 hops 19 should find 17]")

	// test ClosestHops through Shortcuts
	assert.Equal(71, nodes[GetNodeIndexById(5, nodes)].NextClosestHopTo(93).Id, "should find closest hop. [5 hops 93 should find 71]")
	assert.Equal(89, nodes[GetNodeIndexById(22, nodes)].NextClosestHopTo(89).Id, "should find closest hop. [22 hops 89 should find 89 ]")
	assert.Equal(89, nodes[GetNodeIndexById(22, nodes)].NextClosestHopTo(90).Id, "should find closest hop. [22 hops 90 should find 89]")

	//fmt.Println(nodeOne.Lookup(4))

}
