package Slot

import (
	"testing"

	"github.com/HyanSource/Shona/st"
	"github.com/stretchr/testify/assert"
)

/*測試素材管理的方法*/
func TestSymbolManage(t *testing.T) {
	//初始化
	SM := NewSymbolManage()
	//新增
	Add_a := SM.Add(&Symbol{"A", st.NORMAL, []float64{0, 0, 10, 15, 20}})
	Add_k := SM.Add(&Symbol{"K", st.NORMAL, []float64{0, 0, 10, 15, 20}})
	Add_w := SM.Add(&Symbol{"W", st.WILD, []float64{0, 0, 0, 0, 100}})
	Add_s := SM.Add(&Symbol{"S", st.SCATTER, []float64{0, 0, 0, 0, 0}})
	Add_s2 := SM.Add(&Symbol{"S", st.SCATTER, []float64{0, 0, 0, 0, 0}})

	assert.True(t, Add_a)
	assert.True(t, Add_k)
	assert.True(t, Add_w)
	assert.True(t, Add_s)
	assert.False(t, Add_s2)

	//確認
	Check_a := SM.Check("A")
	Check_k := SM.Check("K")
	Check_w := SM.Check("W")
	Check_s := SM.Check("S")
	Check_z := SM.Check("Z")
	assert.True(t, Check_a)
	assert.True(t, Check_k)
	assert.True(t, Check_w)
	assert.True(t, Check_s)
	assert.False(t, Check_z)

	//取得物件

	Get_a := SM.Get("A")
	// Get_k:=SM.Get("K")
	// Get_w:=SM.Get("W")
	// Get_s :=SM.Get("S")
	// Get_z:=SM.Get("Z")

	assert.Equal(t, "A", Get_a.GetName())
	assert.Equal(t, "0", Get_a.GetSymbolType())
	assert.Equal(t, []float64{0, 0, 10, 15, 20}, Get_a.GetOdds())
	//取得賠率

	GetOdds_a := SM.GetOdds("A", 2)

	assert.Equal(t, float64(10), GetOdds_a)

}
