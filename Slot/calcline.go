package Slot

func NewCalcLine(linewidth int) ICalcLine {
	return &CalcLine{}
}

type ICalcLine interface {
	SetLine(line []int) bool
	GetLine(index int) []int
	Len() int
}

type CalcLine struct {
	LineWidth int
	Lines     [][]int
}

func (t *CalcLine) SetLine(line []int) bool {
	return true
}

func (t *CalcLine) GetLine(index int) []int {
	return []int{}
}

func (t *CalcLine) Len() int {
	return 0
}
