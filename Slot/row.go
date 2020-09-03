package Slot

import "errors"

/*初始化行*/
func NewRow(s []string) IRow {
	return &Row{
		symbols: s,
	}
}

type IRow interface {
	SetSymbol(s []string)
	GetLen() int
	GetRow() []string
	GetSymbols(index int, count int) ([]string, error)
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
func (t *Row) GetSymbols(index int, count int) ([]string, error) {

	if count > t.GetLen() {
		return nil, errors.New("count more than the length")
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

	return r, nil
}
