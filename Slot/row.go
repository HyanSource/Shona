package Slot

func NewRoll() IRow {
	return &Row{}
}

type IRow interface {
	SetSymbol()
	GetLength() int
}

//圖案的列
type Row struct {
	symbols []ISymbol //存放列的圖案
}

func (t *Row) SetSymbol() {

}

func (t *Row) GetLength() int {
	return len(t.symbols)
}
