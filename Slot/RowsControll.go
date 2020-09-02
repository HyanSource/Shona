package Slot

/*初始化圖案*/
func NewRowsControll(rs []IRow) IRowsControll {
	return &RowsControll{
		rows: rs,
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
	rows []IRow
}

/*盤面寬度*/
func (t *RowsControll) GetRowsLen() int {
	return len(t.rows)
}

/*取得盤面*/
func (t *RowsControll) GetAllRows() []IRow {
	rs := make([]IRow, t.GetRowsLen())
	copy(rs, t.rows)
	return rs
}

func (t *RowsControll) GetRandomTable() {

}

func (t *RowsControll) GetSpecifyTable() {

}
