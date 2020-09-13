package Slot

import (
	"github.com/HyanSource/Shona/ISlot"
	"github.com/HyanSource/Shona/st"
)

/*初始化圖案*/
func NewSymbol(name string, st st.ST, odds []float64, free []int) ISlot.ISymbol {

	return &Symbol{
		name:       name,
		symboltype: st,
		odds:       odds,
		free:       free,
	}
}

//圖案
type Symbol struct {
	name       string
	symboltype st.ST
	odds       []float64
	free       []int
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

/*取得免費*/
func (t *Symbol) GetFree() []int {
	return t.free
}

/*判斷類型*/
func (t *Symbol) CheckType(st st.ST) bool {
	return t.symboltype == st
}
