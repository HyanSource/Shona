package Slot

/*初始化圖案*/
func NewSymbol(name string, st ST, odds []float64) ISymbol {

	return &Symbol{
		name:       name,
		symboltype: st,
		odds:       odds,
	}
}

type ISymbol interface {
	SetSymbol()
	GetName() string
	GetSymbolType() string
	GetOdds() []float64
	CheckType(st ST) bool
}

//圖案
type Symbol struct {
	name       string
	symboltype ST
	odds       []float64
}

/*重設 節省效能*/
func (t *Symbol) SetSymbol() {

}

/*取得名字*/
func (t *Symbol) GetName() string {
	return t.name
}

/*取得類型*/
func (t *Symbol) GetSymbolType() string {
	return t.symboltype.ToString()
}

/*取得賠率*/
func (t *Symbol) GetOdds() []float64 {
	return t.odds
}

/*判斷類型*/
func (t *Symbol) CheckType(st ST) bool {
	return t.symboltype == st
}
