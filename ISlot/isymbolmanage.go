package ISlot

type ISymbolManage interface {
	Add(s ISymbol) bool                     //新增元素
	Check(s string) bool                    //檢查元素
	Get(s string) ISymbol                   //取得元素
	GetOdds(s string, linelen int) float64  //取得賠率
	GetFreeCount(s string, linelen int) int //取得免費次數
	RemoveAll()                             //移除所有
}
