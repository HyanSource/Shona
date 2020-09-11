package Slot

import (
	"fmt"
	"testing"

	"github.com/HyanSource/Shona/st"

	"github.com/HyanSource/Shona/ISlot"
)

func TestCalcline(t *testing.T) {

	//https://ezslotdesign.com/slotmath1v2/

	//滾輪條
	rs := []ISlot.IRow{
		NewRow([]string{}),
		NewRow([]string{}),
		NewRow([]string{}),
		NewRow([]string{}),
		NewRow([]string{}),
	}

	fmt.Println(rs)

	sm := NewSymbolManage()
	sm.Add(NewSymbol("K", st.NORMAL, []float64{0, 0, 5, 30, 100}))
	sm.Add(NewSymbol("Q", st.NORMAL, []float64{0, 0, 5, 25, 100}))
	sm.Add(NewSymbol("J", st.NORMAL, []float64{0, 0, 5, 20, 75}))
	sm.Add(NewSymbol("10", st.NORMAL, []float64{0, 0, 5, 20, 75}))
	sm.Add(NewSymbol("9", st.NORMAL, []float64{0, 0, 5, 10, 50}))
	sm.Add(NewSymbol("S", st.SCATTER, []float64{0, 0, 50, 0, 0}))
}
