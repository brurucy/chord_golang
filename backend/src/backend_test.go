package src

import (
	"fmt"
	"testing"
)

func TestRingDistance(t *testing.T) {

	from := []int{22, 4, 4, 22, 4, 4}
	to := []int{4, 22, 4, 4, 22, 4}
	maxSize := []int{100, 100, 100, 100, 100, 100}
	minSize := []int{0, 0, 0, 1, 1, 1}
	comparisons := []int{82, 18, 0, 81, 18, 0}

	var testOne int

	for i := 0; i < 6; i++ {

		testOne = RingDistance(from[i], to[i], maxSize[i], minSize[i])

		if testOne != comparisons[i] {

			t.Error("frzom: ", from[i], "to: ", to[i], "dist: ", testOne, " != ", comparisons[i])

		}

		fmt.Println("Dist: ", testOne)

	}

}

func TestStabilizeAddShortcut(t *testing.T) {

	nodeOne := Node{
		Id: 5,
	}

	nodeTwo := Node{
		Id: 17,
	}

	nodeThree := Node{
		Id: 22,
	}

	nodeFour := Node{
		Id: 56,
	}

	nodeFive := Node{
		Id: 71,
	}

	nodeSix := Node{
		Id: 89,
	}

	nodeSeven := Node{
		Id: 92,
	}

	//nodeEight := Node{
	//	Id: 110,
	//}
	//
	const maxrange = 100

	nodeOne.Succ = &nodeTwo
	nodeTwo.Succ = &nodeThree
	nodeThree.Succ = &nodeFour
	nodeFour.Succ = &nodeFive
	nodeFive.Succ = &nodeSix
	nodeSix.Succ = &nodeSeven
	nodeSeven.Succ = &nodeOne
	//nodeEight.Succ = &nodeOne

	fmt.Println(nodeOne)

	nodeOne.Stabilize()
	nodeTwo.Stabilize()
	nodeThree.Stabilize()
	nodeFour.Stabilize()
	nodeFive.Stabilize()
	nodeSix.Stabilize()
	nodeSeven.Stabilize()

	nodeOne.AddShortcut(&nodeFour)
	nodeOne.AddShortcut(&nodeFive)
	nodeThree.AddShortcut(&nodeSix)

	nodeOne.MigrateData(maxrange)
	nodeTwo.MigrateData(maxrange)

	/*nodeThree.MigrateData(maxrange)
	nodeFour.MigrateData(maxrange)
	nodeFive.MigrateData(maxrange)
	nodeSix.MigrateData(maxrange)
	nodeSeven.MigrateData(maxrange)
	*/

	if &nodeOne != nodeSeven.Succ {

		t.Errorf("failed to stabilize")

	}

	if len(nodeOne.Shortcuts) != 2 {

		t.Errorf("failed to add shortcuts")

	}

	//fmt.Println(nodeTwo.ClosestHopTo(93))

	//fmt.Println(nodeOne.findValue(4))

}
