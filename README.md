# Shona
輕量級Slot框架

#### examples
```go
func main() {
//新增5行滾輪條
	rs := []ISlot.IRow{
		Slot.NewRow([]string{"K", "Q", "J", "10", "9", "9", "K", "Q", "J", "10", "9", "10", "Q", "J", "10", "9", "9", "J", "10", "9"}),
		Slot.NewRow([]string{"W", "K", "Q", "J", "10", "9", "S", "K", "Q", "J", "10", "9", "K", "Q", "J", "10", "9", "J", "J", "9"}),
		Slot.NewRow([]string{"W", "K", "Q", "J", "10", "9", "S", "K", "Q", "J", "10", "9", "K", "Q", "J", "10", "9", "J", "10", "9"}),
		Slot.NewRow([]string{"W", "K", "Q", "J", "10", "9", "S", "K", "Q", "J", "10", "9", "K", "Q", "J", "10", "9", "J", "10", "9"}),
		Slot.NewRow([]string{"W", "K", "Q", "J", "10", "9", "9", "K", "Q", "J", "10", "9", "K", "Q", "J", "10", "9", "J", "10", "9"}),
	}

	//新增素材類型以及賠率
	sm := Slot.NewSymbolManage()
	sm.Add(Slot.NewSymbol("K", st.NORMAL, []float64{0, 0, 5, 30, 100}))
	sm.Add(Slot.NewSymbol("Q", st.NORMAL, []float64{0, 0, 5, 25, 100}))
	sm.Add(Slot.NewSymbol("J", st.NORMAL, []float64{0, 0, 5, 20, 75}))
	sm.Add(Slot.NewSymbol("10", st.NORMAL, []float64{0, 0, 5, 20, 75}))
	sm.Add(Slot.NewSymbol("9", st.NORMAL, []float64{0, 0, 5, 10, 50}))
	sm.Add(Slot.NewSymbol("S", st.SCATTER, []float64{0, 0, 50, 0, 0}))
	sm.Add(Slot.NewSymbol("W", st.WILD, []float64{0, 0, 0, 0, 0}))

	//新增LineGame計算 設置高度 滾輪 素材管理
	cl := Slot.NewCalcLine(3, rs, sm)
	cl.Init()

	result := cl.GetResult()
	fmt.Println("RTP:", result.GetRTP())
	fmt.Println("FreeP:", result.GetFreeP())
}
```

#### 架構圖
![image]("https://github.com/HyanSource/Shona/blob/master/ShonaArchitecture.png")