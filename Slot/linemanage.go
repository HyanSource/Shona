package Slot

import "github.com/HyanSource/Shona/ISlot"

/*
 * 設置陣列數字規則
 * ex: 5*3 slot
 * .  .  .  .  .
 * 10 11 12 13 14
 * 5  6  7  8  9
 * 0  1  2  3  4
 */

func NewLineManage(linewidth int, height int) ISlot.ILineManage {
	return &LineManage{
		linewidth:  linewidth,
		lineheight: height,
	}
}

/*線的規則*/
type LineManage struct {
	linewidth  int     //滾輪寬度 (判斷有無超過寬度)
	lineheight int     //滾輪高度 (判斷數字有無超過最大)
	lines      [][]int //手動加入
}

/*設置一行中獎線 必須和滾輪長度符合*/
func (t *LineManage) SetLine(line []int) bool {

	if len(line) == t.linewidth {
		t.lines = append(t.lines, line)
		return true
	}
	return false
}

/*取得一條中獎線*/
func (t *LineManage) GetLine(index int) []int {
	return t.lines[index]
}

/*清除所有*/
func (t *LineManage) RemoveAllLines() {
	t.lines = make([][]int, 0)
}

/*取得中獎線的數量*/
func (t *LineManage) Len() int {
	return len(t.lines)
}
