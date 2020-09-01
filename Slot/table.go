package Slot

/*初始化圖案*/
func NewTable(rs []IRow) ITable {
	return &Table{
		rows: rs,
	}
}

type ITable interface {
	GetTableLen() int
	GetTable() []IRow
	GetWildCount() int
	GetScatterCount() int
	GetOdds() float64

	GetRandomTable()
	GetSpecifyTable()
}

/*盤面*/
type Table struct {
	rows []IRow
}

/*盤面寬度*/
func (t *Table) GetTableLen() int {
	return len(t.rows)
}

/*取得盤面*/
func (t *Table) GetTable() []IRow {
	rs := make([]IRow, t.GetTableLen())
	copy(rs, t.rows)
	return rs
}

/*盤面wild數量*/
func (t *Table) GetWildCount() int {
	count := 0
	for _, v := range t.rows {
		count += v.GetWildCount()
	}
	return count
}

func (t *Table) GetScatterCount() int {
	return 0
}

func (t *Table) GetOdds() float64 {
	return 0
}

func (t *Table) GetRandomTable() {

}

func (t *Table) GetSpecifyTable() {

}
