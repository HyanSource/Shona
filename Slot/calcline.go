package Slot

import (
	"sort"
	"strconv"

	"github.com/HyanSource/Shona/ISlot"
)

func NewCalcLine(rs []ISlot.IRow, sm ISlot.ISymbolManage) ISlot.ICalc {
	return &CalcLine{
		rows:         rs,
		SymbolManage: sm,
	}
}

type CalcLine struct {
	rows         []ISlot.IRow        //所有盤面滾輪
	SymbolManage ISlot.ISymbolManage //素材管理

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

func (t *CalcLine) GetPossibility() int {
	return t.possibility
}

func (t *CalcLine) GetGameTable(index int) []string {

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

	return []string{}
}

/*計算賠率部分*/
func (t *CalcLine) CalcOdds(result []string) float64 {
	lott := result[0]
	count := 1
	for i := 1; i < len(result); i++ {
		/*等於自身符號次或是WILD*/
		if result[i] == lott || t.SymbolManage.Get(result[i]).GetSymbolType() == "1" {
			count++
		} else {
			break
		}
	}

	return t.SymbolManage.GetOdds(lott, count)
}

/*計算scatter賠率部分和獲得免費次數(依照不同規則)*/
func (t *CalcLine) CalcScatter(result []string) (float64, int) {

	/*目前算只有考慮當只有1種Scatter*/
	count := 0
	scatter := ""
	for i := 0; i < len(result); i++ {
		if t.SymbolManage.Get(result[i]).GetSymbolType() == "2" {
			if count == 0 {
				scatter = t.SymbolManage.Get(result[i]).GetName()
				count++
			} else if scatter == t.SymbolManage.Get(result[i]).GetName() {
				count++
			}
		}
	}
	/*暫時寫法*/
	return t.SymbolManage.GetOdds("S", count), 0
}

/*總共RTP*/
func (t *CalcLine) GetRTP() float64 {

	for i := 0; i < t.possibility; i++ {

	}

	return 0
}

/*取得免費觸發機率*/
func (t *CalcLine) GetFreeP() float64 {

	for i := 0; i < t.possibility; i++ {

	}

	return 0
}

//模擬正常玩的RTP
func (t *CalcLine) GetRandomPlayRTP(playcount int) float64 {

	for i := 0; i < playcount; i++ {

	}

	return 0
}

/*需要一個result的interface*/
