package Slot

/*初始化行*/
func NewRow(s []ISymbol) IRow {
	return &Row{
		symbols: s,
	}
}

type IRow interface {
	SetSymbol(s []ISymbol)
	GetLen() int
	GetRow() []ISymbol
	GetWildCount() int
}

//圖案的列
type Row struct {
	symbols []ISymbol //存放列的圖案
}

/*重設 節省效能*/
func (t *Row) SetSymbol(s []ISymbol) {
	t.symbols = s
}

/*取得行長度*/
func (t *Row) GetLen() int {
	return len(t.symbols)
}

/*取得行物件*/
func (t *Row) GetRow() []ISymbol {
	r := make([]ISymbol, t.GetLen())
	copy(r, t.symbols)
	return r
}

/*一行的wild數量*/
func (t *Row) GetWildCount() int {
	count := 0
	for _, v := range t.symbols {
		if v.CheckWild() {
			count++
		}
	}
	return count
}
