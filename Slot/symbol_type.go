package Slot

import "fmt"

//symbol type
type ST int

const (
	NORMAL  ST = iota //0
	WILD              //1
	SCATTER           //2
)

//to string
func (t ST) ToString() string {
	switch t {
	case NORMAL, WILD, SCATTER:
		return fmt.Sprintf("%v", t)
	default:
		panic(fmt.Sprintf("not exists type: %v", t))
	}
}
