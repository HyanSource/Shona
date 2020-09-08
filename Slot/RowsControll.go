package Slot

/*初始化圖案*/
func NewRowsControll(rs []IRow, sm ISymbolManage) IRowsControll {
	return &RowsControll{
		rows:         rs,
		SymbolManage: sm,
	}
}

type IRowsControll interface {
	GetRowsLen() int
	GetAllRows() []IRow

	GetRandomTable()
	GetSpecifyTable()
}

/*盤面*/
type RowsControll struct {
	rows         []IRow        //所有盤面滾輪
	SymbolManage ISymbolManage //素材管理
	calcline     ICalcLine     //線規則
}

/*取得盤面寬度*/
func (t *RowsControll) GetRowsLen() int {
	return len(t.rows)
}

/*取得所有盤面*/
func (t *RowsControll) GetAllRows() []IRow {
	rs := make([]IRow, t.GetRowsLen())
	copy(rs, t.rows)
	return rs
}

func (t *RowsControll) GetRandomTable() {

}

func (t *RowsControll) GetSpecifyTable() {

}
