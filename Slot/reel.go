package Slot

import "github.com/HyanSource/Shona/ISlot"

/*初始化行*/
func Newreel(s []string) ISlot.Ireel {
	return &reel{
		symbols: s,
	}
}

//圖案的列
type reel struct {
	symbols []string //存放列的圖案
}

/*重設 節省效能*/
func (t *reel) SetSymbol(s []string) {
	t.symbols = s
}

/*取得行長度*/
func (t *reel) GetLen() int {
	return len(t.symbols)
}

/*取得行物件*/
func (t *reel) Getreel() []string {
	r := make([]string, t.GetLen())
	copy(r, t.symbols)
	return r
}

/**/
func (t *reel) GetSymbols(index int, count int) []string {

	if count > t.GetLen() {
		return nil
	}

	r := make([]string, count)

	for i := 0; i < count; i++ {
		r[i] = t.symbols[index]

		if index == t.GetLen() {
			index = 0
		} else {
			index++
		}
	}

	return r
}
