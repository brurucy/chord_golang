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
	nodes[0].AddShortcut(&nodes[3])
	nodes[0].AddShortcut(&nodes[4])
	nodes[2].AddShortcut(&nodes[5])

	// migrate data TODO: for all nodes?
	for i := 0; i < 2; /*len(nodeIds)*/ i++ {
		nodes[i].MigrateData(maxrange)
	}

	assert.Equal(nodes[len(nodes)-1].Succ, &nodes[0], "last node doesn't link to first")
	assert.Equal(nodes[0].Succ.Succ, &nodes[2], "failed to stabilize SuccSucc")
	assert.Equal(len(nodes[0].Shortcuts), 2, "failed to add shortcuts")

	//fmt.Println(nodeTwo.ClosestHopTo(93))

	//fmt.Println(nodeOne.findValue(4))

}
