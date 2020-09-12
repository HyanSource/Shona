package Slot

import "github.com/HyanSource/Shona/ISlot"

func NewSymbolManage() ISlot.ISymbolManage {
	return &SymbolManage{
		symbols: make(map[string]ISlot.ISymbol),
	}
}

//管理所有的素材和賠率
type SymbolManage struct {
	symbols map[string]ISlot.ISymbol
}

/*加入素材*/
func (t *SymbolManage) Add(s ISlot.ISymbol) bool {

	if !t.Check(s.GetName()) {
		t.symbols[s.GetName()] = s
		return true
	}
	return false
}

/*確認有無素材*/
func (t *SymbolManage) Check(s string) bool {
	_, ok := t.symbols[s]
	return ok
}

/*取得素材*/
func (t *SymbolManage) Get(s string) ISlot.ISymbol {

	if t.Check(s) {
		return t.symbols[s]
	}

	return nil
}

/*取得賠率*/
func (t *SymbolManage) GetOdds(s string, linelen int) float64 {

	if t.Check(s) && linelen > 0 {
		return t.Get(s).GetOdds()[linelen-1]
	}
	return 0
}

/*移除全部*/
func (t *SymbolManage) RemoveAll() {
	t.symbols = make(map[string]ISlot.ISymbol)
}
