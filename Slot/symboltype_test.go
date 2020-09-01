package Slot

import (
	"fmt"
	"testing"
)

func TestSymbolType(t *testing.T) {

	//first way
	var t1 ST
	t1 = NORMAL
	fmt.Println(t1, " ", t1.ToString())

	//second way
	t2 := ST(0)
	fmt.Println(t2, " ", t2.ToString())

	//third way
	var t3 ST
	t3 = 2
	fmt.Println(t3, " ", t3.ToString())
}
