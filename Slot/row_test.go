package Slot

import (
	"fmt"
	"testing"
)

func TestRow(t *testing.T) {
	a := NewSymbol("A", NORMAL, []int{0, 0, 10, 15, 20})
	fmt.Println(a)
}
