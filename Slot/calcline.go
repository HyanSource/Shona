package Slot

/*
 * 設置陣列數字規則
 * ex: 5*3 slot
 * .  .  .  .  .
 * 10 11 12 13 14
 * 5  6  7  8  9
 * 0  1  2  3  4
 */

func NewCalcLine(linewidth int, height int) ICalcLine {
	return &CalcLine{
		linewidth:  linewidth,
		lineheight: height,
	}
}

type ICalcLine interface {
	SetLine(line []int) bool
	GetLine(index int) []int
	RemoveAllLines()
	Len() int
}

/*線的規則*/
type CalcLine struct {
	linewidth  int     //滾輪寬度 由RowsControll設置
	lineheight int     //滾輪高度 由RowsControll設置
	lines      [][]int //手動加入
}

/*設置一行中獎線 必須和滾輪長度符合*/
func (t *CalcLine) SetLine(line []int) bool {

	if len(line) == t.linewidth {
		t.lines = append(t.lines, line)
		return true
	}
	return false
}

/*取得一條中獎線*/
func (t *CalcLine) GetLine(index int) []int {
	return t.lines[index]
}

/*清除所有*/
func (t *CalcLine) RemoveAllLines() {
	t.lines = make([][]int, 0)
}

/*取得中獎線的數量*/
func (t *CalcLine) Len() int {
	return len(t.lines)
}
