package Slot

//初始化圖案
func NewSymbol(name string, typeint int, length int) ISymbol {

	return &Symbol{
		name:       name,
		symboltype: ST(typeint),
		odds:       make([]int, length),
	}
}

type ISymbol interface {
	SetSymbol()
	GetName() string
	GetSymbolType() string
	GetOdds() []int
}

//圖案
type Symbol struct {
	name       string
	symboltype ST
	odds       []int
}

func (t *Symbol) SetSymbol() {

}

func (t *Symbol) GetName() string {
	return t.name
}

func (t *Symbol) GetSymbolType() string {
	return t.symboltype.ToString()
}

func (t *Symbol) GetOdds() []int {
	return t.odds
}
