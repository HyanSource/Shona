package Slot

func NewSymbolManage() ISymbolManage {
	return &SymbolManage{
		symbols: make(map[string]ISymbol),
	}
}

type ISymbolManage interface {
	Add(s ISymbol) bool                    //新增元素
	Check(s string) bool                   //檢查元素
	Get(s string) ISymbol                  //取得元素
	GetOdds(s string, linelen int) float64 //取得賠率
	RemoveAll()                            //移除所有
}

//管理所有的素材和賠率
type SymbolManage struct {
	symbols map[string]ISymbol
}

/*加入素材*/
func (t *SymbolManage) Add(s ISymbol) bool {

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
func (t *SymbolManage) Get(s string) ISymbol {

	if t.Check(s) {
		return t.symbols[s]
	}

	return nil
}

/*取得賠率*/
func (t *SymbolManage) GetOdds(s string, linelen int) float64 {
	if t.Check(s) {
		return t.Get(s).GetOdds()[linelen]
	}
	return 0
}

/*移除全部*/
func (t *SymbolManage) RemoveAll() {
	t.symbols = make(map[string]ISymbol)
}
