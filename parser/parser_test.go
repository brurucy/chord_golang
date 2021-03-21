package parser

import (
	"fmt"
	"testing"
)

func TestParser(t *testing.T) {

	ble, _ := parse("../Input-file.txt")

	fmt.Println(ble)

}