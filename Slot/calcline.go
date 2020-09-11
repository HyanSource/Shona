package Slot

import (
	"sort"
	"strconv"

	"github.com/HyanSource/Shona/ISlot"
)

func NewCalcLine(height int, rs []ISlot.IRow, sm ISlot.ISymbolManage) ISlot.ICalc {
	return &CalcLine{
		rowsheight:   height,
		rows:         rs,
		SymbolManage: sm,
	}
}

type CalcLine struct {
	rowsheight   int                 //盤面滾輪高度
	rows         []ISlot.IRow        //所有盤面滾輪
	SymbolManage ISlot.ISymbolManage //素材管理
	LineManage   ISlot.ILineManage   //線規則物件

	possibility int   //所有可能性
	subarray    []int //用來扣除來算出相對應的盤面陣列
}

func (t *CalcLine) Init() {
	t.possibility = 1
	t.subarray = make([]int, len(t.rows))
	for k, v := range t.rows {
		t.subarray[k] = t.possibility
		t.possibility *= v.GetLen()
	}

	sort.Sort(sort.Reverse(sort.IntSlice(t.subarray)))
}

func (t *CalcLine) GetLen() int {
	return len(t.rows)
}
func (t *CalcLine) GetHeight() int {
	return t.rowsheight
}
func (t *CalcLine) GetPossibility() int {
	return t.possibility
}

func (t *CalcLine) GetGameTable(index int) ([]string, float64) {

	if index >= t.GetPossibility() || index < 0 {
		panic("no table possibility:" + strconv.Itoa(index))
	}
	//滾輪引索
	index_arr := make([]int, len(t.rows))
	for i := 0; i < len(index_arr); i++ {
		index_arr[i] = (index / t.subarray[i])
		index = (index % t.subarray[i])
	}

	//顯示盤面
	result := make([]string, len(t.rows))
	for i := 0; i < len(result); i++ {
		result[i] = t.rows[i].GetSymbols(index_arr[i], 1)[0]
	}

	return []string{}, 0
}

func (t *CalcLine) GetRTP() float64 {
	return 0
}
