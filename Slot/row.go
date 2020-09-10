package Slot

import "github.com/HyanSource/Shona/ISlot"

/*初始化行*/
func NewRow(s []string) ISlot.IRow {
	return &Row{
		symbols: s,
	}
}

//圖案的列
type Row struct {
	symbols []string //存放列的圖案
}

/*重設 節省效能*/
func (t *Row) SetSymbol(s []string) {
	t.symbols = s
}

/*取得行長度*/
func (t *Row) GetLen() int {
	return len(t.symbols)
}

/*取得行物件*/
func (t *Row) GetRow() []string {
	r := make([]string, t.GetLen())
	copy(r, t.symbols)
	return r
}

/**/
func (t *Row) GetSymbols(index int, count int) []string {

	if count > t.GetLen() {
		return nil
	}

	r := make([]string, count)

	for i := 0; i < count; i++ {
		r[i] = t.symbols[index]

		if index == t.GetLen() {
			index = 0
		} else {
			index++
		}
	}

	return r
}
