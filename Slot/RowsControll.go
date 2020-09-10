package Slot

import (
	"strconv"

	"github.com/HyanSource/Shona/ISlot"
	_ "github.com/HyanSource/Shona/ISlot"
)

/*初始化圖案*/
func NewRowsControll(rs []ISlot.IRow, sm ISlot.ISymbolManage, rowsheight int) ISlot.IRowsControll {
	return &RowsControll{
		rows:         rs,
		SymbolManage: sm,
		rowsheight:   rowsheight,
		LineManage:   NewLineManage(len(rs), rowsheight),
	}
}

/*盤面*/
type RowsControll struct {
	rowsheight   int                 //盤面滾輪高度
	rows         []ISlot.IRow        //所有盤面滾輪
	SymbolManage ISlot.ISymbolManage //素材管理
	LineManage   ISlot.ILineManage   //線規則物件
}

/*取得盤面寬度*/
func (t *RowsControll) GetRowsLen() int {
	return len(t.rows)
}

/*取得所有盤面*/
func (t *RowsControll) GetAllRows() []ISlot.IRow {
	rs := make([]ISlot.IRow, t.GetRowsLen())
	copy(rs, t.rows)
	return rs
}

//取得所有機率發生次數
func (t *RowsControll) GetPossibility() int {
	p := 1
	for _, v := range t.rows {
		p *= v.GetLen()
	}
	return p
}

//取得特定盤面
func (t *RowsControll) GetGameTable(index int) ([]string, float64) {

	if index >= t.GetPossibility() || index < 0 {
		panic("no table possibility:" + strconv.Itoa(index))
	}

	return []string{}, 0
}

//取得RTP
func (t *RowsControll) GetLineGameRTP() float64 {

	return 0
}

//取得RTP
func (t *RowsControll) GetWayGameRTP() float64 {
	return 0
}

//可能需要新增兩個東西 calcline calcway
//丟進滾輪條(pointer) 產生盤面
//把linemanage放進去
//方法有 index產生盤面 計算賠率

//更進階的寫法 定義ICalc讓2個規則去實作
