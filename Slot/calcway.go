package Slot

import (
	"sort"
	"strconv"

	"github.com/HyanSource/Shona/ISlot"
)

func NewCalcWay(height int, rs []ISlot.Ireel, sm ISlot.ISymbolManage) ISlot.ICalc {
	return &CalcWay{
		height:       height,
		reels:        rs,
		SymbolManage: sm,
	}
}

type CalcWay struct {
	height       int
	reels        []ISlot.Ireel
	SymbolManage ISlot.ISymbolManage

	possibility int
	subarray    []int
}

func (t *CalcWay) Init() {
	t.possibility = 1
	t.subarray = make([]int, len(t.reels))
	for k, v := range t.reels {
		t.subarray[k] = t.possibility
		t.possibility *= v.GetLen()
	}

	sort.Sort(sort.Reverse(sort.IntSlice(t.subarray))) //反轉陣列
}

func (t *CalcWay) GetLen() int {
	return len(t.reels)
}

func (t *CalcWay) GetPossibility() int {
	return t.possibility
}

func (t *CalcWay) GetGameTable(index int) []string {

	if index >= t.GetPossibility() || index < 0 {
		panic("no table possibility:" + strconv.Itoa(index))
	}

	//滾輪引索
	index_arr := make([]int, len(t.reels))
	for i := 0; i < len(t.reels); i++ {
		index_arr[i] = (index / t.subarray[i])
		index %= t.subarray[i]
	}

	//顯示盤面
	result := make([]string, 0, len(t.reels))
	for i := 0; i < len(t.reels); i++ {
		m := t.reels[i].GetSymbols(index_arr[i], 3)
		for j := 0; j < len(m); j++ {
			result = append(result, m[j])
		}
	}

	return result
}

func (t *CalcWay) CalcOdds(result []string) float64 {

	// multi := 0
	// count := 1

	//2  5  8  11 14
	//1  4  7  10 13
	//0  3  6  9  12
	for i := 1; i < len(result); i++ {

	}

	return 0.0
}

func (t *CalcWay) CalcScatter(result []string) (float64, int, int) {
	return 0.0, 0, 0
}

func (t *CalcWay) GetResult() ISlot.IResult {
	return &Result{}
}

func (t *CalcWay) GetNormalPlayResult(playcount int) ISlot.IResult {
	return &Result{}
}
