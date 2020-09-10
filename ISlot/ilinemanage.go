package ISlot

type ILineManage interface {
	SetLine(line []int) bool
	GetLine(index int) []int
	RemoveAllLines()
	Len() int
}
