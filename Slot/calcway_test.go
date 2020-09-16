package Slot

import (
	"fmt"
	"testing"

	"github.com/HyanSource/Shona/ISlot"
	"github.com/HyanSource/Shona/st"
	"github.com/stretchr/testify/assert"
)

func TestCalcWay(t *testing.T) {
	//滾輪條
	rs := []ISlot.Ireel{
		Newreel([]string{"K", "Q", "J", "10", "9", "9", "K", "Q", "J", "10", "9", "10", "Q", "J", "10", "9", "9", "J", "10", "9"}),
		Newreel([]string{"W", "K", "Q", "J", "10", "9", "S", "K", "Q", "J", "10", "9", "K", "Q", "J", "10", "9", "J", "J", "9"}),
		Newreel([]string{"W", "K", "Q", "J", "10", "9", "S", "K", "Q", "J", "10", "9", "K", "Q", "J", "10", "9", "J", "10", "9"}),
		Newreel([]string{"W", "K", "Q", "J", "10", "9", "S", "K", "Q", "J", "10", "9", "K", "Q", "J", "10", "9", "J", "10", "9"}),
		Newreel([]string{"W", "K", "Q", "J", "10", "9", "9", "K", "Q", "J", "10", "9", "K", "Q", "J", "10", "9", "J", "10", "9"}),
	}

	sm := NewSymbolManage()
	assert.True(t, sm.Add(NewSymbol("K", st.NORMAL, []float64{0, 0, 5, 30, 100}, []int{})))
	assert.True(t, sm.Add(NewSymbol("Q", st.NORMAL, []float64{0, 0, 5, 25, 100}, []int{})))
	assert.True(t, sm.Add(NewSymbol("J", st.NORMAL, []float64{0, 0, 5, 20, 75}, []int{})))
	assert.True(t, sm.Add(NewSymbol("10", st.NORMAL, []float64{0, 0, 5, 20, 75}, []int{})))
	assert.True(t, sm.Add(NewSymbol("9", st.NORMAL, []float64{0, 0, 5, 10, 50}, []int{})))
	assert.True(t, sm.Add(NewSymbol("S", st.SCATTER, []float64{0, 0, 50, 0, 0}, []int{0, 0, 10, 15, 20})))
	assert.True(t, sm.Add(NewSymbol("W", st.WILD, []float64{0, 0, 0, 0, 0}, []int{})))

	cl := NewCalcWay(3, rs, sm)
	cl.Init()

	// assert.Equal(t, 5, cl.GetLen())
	// length := rs[0].GetLen() * rs[1].GetLen() * rs[2].GetLen() * rs[3].GetLen() * rs[4].GetLen()
	// assert.Equal(t, length, cl.GetPossibility())

	// assert.Equal(t, float64(100), cl.CalcOdds([]string{"K", "K", "K", "K", "K"}))
	// assert.Equal(t, float64(100), cl.CalcOdds([]string{"K", "W", "K", "K", "K"}))
	// assert.Equal(t, float64(100), cl.CalcOdds([]string{"K", "K", "W", "K", "K"}))
	// assert.Equal(t, float64(100), cl.CalcOdds([]string{"K", "K", "K", "W", "K"}))
	// assert.Equal(t, float64(100), cl.CalcOdds([]string{"K", "K", "K", "K", "W"}))

	// assert.Equal(t, float64(30), cl.CalcOdds([]string{"K", "K", "K", "K", "Q"}))
	// assert.Equal(t, float64(30), cl.CalcOdds([]string{"K", "W", "K", "K", "Q"}))
	// assert.Equal(t, float64(30), cl.CalcOdds([]string{"K", "K", "W", "K", "Q"}))
	// assert.Equal(t, float64(30), cl.CalcOdds([]string{"K", "K", "K", "W", "Q"}))

	// assert.Equal(t, float64(5), cl.CalcOdds([]string{"K", "K", "K", "10", "Q"}))
	// assert.Equal(t, float64(5), cl.CalcOdds([]string{"K", "W", "K", "9", "Q"}))
	// assert.Equal(t, float64(5), cl.CalcOdds([]string{"K", "K", "W", "S", "Q"}))
	// odds, _, _ := cl.CalcScatter([]string{"K", "S", "S", "S", "Q"})
	// assert.Equal(t, float64(50), odds)

	fmt.Println(cl.GetGameTable(0))
}
