package Slot

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/HyanSource/Shona/st"

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
		m := t.reels[i].GetSymbols(index_arr[i], t.height)
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

	//每個滾輪上創建map 當沒有元素就創建 有則是增加count
	lotts := make([]map[string]int, len(t.reels))

	for i := 0; i < len(t.reels); i++ {
		for j := 0; j < t.height; j++ {
			value, ok := lotts[i][result[0]]
			if ok {
				value++
			}
			lotts[i][result[0]] = 1
		}
	}

	fmt.Println(lotts)
	//第一格滾輪
	for k, v := range lotts[0] {
		lott := k
		odds := v
		for j := 1; j < len(lotts); j++ {
			value, ok := lotts[j][lott]
			if ok {
				odds *= value
			} else {
				break
			}
		}
	}

	return 0.0
}

func (t *CalcWay) CalcScatter(result []string) (float64, int, int) {

	/**/

	count := 0
	scatter := ""
	for i := 0; i < len(result); i++ {
		if t.SymbolManage.Get(result[i]).CheckType(st.SCATTER) {
			if count == 0 {
				scatter = t.SymbolManage.Get(result[i]).GetName()
				count++
			} else if scatter == t.SymbolManage.Get(result[i]).GetName() {
				count++
			}
		}

	}

	return t.SymbolManage.GetOdds("S", count), count, t.SymbolManage.GetFreeCount("S", count)
}

func (t *CalcWay) GetResult() ISlot.IResult {
	return &Result{}
}

func (t *CalcWay) GetNormalPlayResult(playcount int) ISlot.IResult {
	return &Result{}
}
