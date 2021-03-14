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

	const maxrange = 100
	nodeIds := []int{5, 17, 22, 56, 71, 89, 92} // 110
	var nodes []Node

	// create nodes
	for i := 0; i < len(nodeIds); i++ {
		nodes = append(nodes, Node{Id: nodeIds[i]})
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
	// nodes[GetNodeIndexById(92, nodes)].AddShortcut(&nodes[GetNodeIndexById(56, nodes)]) // TODO: fix stack overflow

	// migrate data TODO: for all nodes?
	for i := 0; i < 2; /*len(nodeIds)*/ i++ {
		nodes[i].MigrateData(maxrange)
	}

	// test shortcuts adding
	assert.Equal(len(nodes[0].Shortcuts), 2, "should add shortcuts")

	// test stabilization
	assert.Equal(nodes[len(nodes)-1].Succ, &nodes[0], "last node in ring should have first node as Succ")
	assert.Equal(nodes[0].Succ.Succ, &nodes[2], "should find and add SuccSucc after stabilize")

	// test ClosestHops through SuccSucc
	assert.Equal(56, nodes[GetNodeIndexById(22, nodes)].ClosestHopTo(56).Id, "should find closest hop [Succ, target, later in circle]")
	assert.Equal(56, nodes[GetNodeIndexById(22, nodes)].ClosestHopTo(50).Id, "should find closest hop [Succ, target, later in circle]")
	assert.Equal(5, nodes[GetNodeIndexById(92, nodes)].ClosestHopTo(4).Id, "should find closest hop [Succ, target, accross circle start]")
	assert.Equal(5, nodes[GetNodeIndexById(92, nodes)].ClosestHopTo(93).Id, "should find closest hop [Succ, target, accross circle start]")

	// test ClosestHops through SuccSucc
	assert.Equal(17, nodes[GetNodeIndexById(92, nodes)].ClosestHopTo(10).Id, "should find closest hop [SuccSucc, target, across circle start]")
	assert.Equal(17, nodes[GetNodeIndexById(92, nodes)].ClosestHopTo(19).Id, "should find closest hop [SuccSucc, before target, across circle start]")

	// test ClosestHops through Shortcuts
	assert.Equal(71, nodes[GetNodeIndexById(5, nodes)].ClosestHopTo(93).Id, "should find closest hop [Shortcut, before target, later in circle]")
	assert.Equal(89, nodes[GetNodeIndexById(22, nodes)].ClosestHopTo(89).Id, "should find closest hop [Shortcut, target, later in circle]")
	assert.Equal(56, nodes[GetNodeIndexById(89, nodes)].ClosestHopTo(55).Id, "should find closest hop [Shortcut, target, across circle start]")

	//fmt.Println(nodeOne.findValue(4))

}
