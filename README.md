# Shona
輕量級Slot計算RTP框架

#### examples
```go
func main() {
//新增5行滾輪條
	rs := []ISlot.Ireel{
		Slot.Newreel([]string{"K", "Q", "J", "10", "9", "9", "K", "Q", "J", "10", "9", "10", "Q", "J", "10", "9", "9", "J", "10", "9"}),
		Slot.Newreel([]string{"W", "K", "Q", "J", "10", "9", "S", "K", "Q", "J", "10", "9", "K", "Q", "J", "10", "9", "J", "J", "9"}),
		Slot.Newreel([]string{"W", "K", "Q", "J", "10", "9", "S", "K", "Q", "J", "10", "9", "K", "Q", "J", "10", "9", "J", "10", "9"}),
		Slot.Newreel([]string{"W", "K", "Q", "J", "10", "9", "S", "K", "Q", "J", "10", "9", "K", "Q", "J", "10", "9", "J", "10", "9"}),
		Slot.Newreel([]string{"W", "K", "Q", "J", "10", "9", "9", "K", "Q", "J", "10", "9", "K", "Q", "J", "10", "9", "J", "10", "9"}),
	}

	//新增素材類型以及賠率
	sm := Slot.NewSymbolManage()
	sm.Add(Slot.NewSymbol("K", st.NORMAL, []float64{0, 0, 5, 30, 100}, []int{}))
	sm.Add(Slot.NewSymbol("Q", st.NORMAL, []float64{0, 0, 5, 25, 100}, []int{}))
	sm.Add(Slot.NewSymbol("J", st.NORMAL, []float64{0, 0, 5, 20, 75}, []int{}))
	sm.Add(Slot.NewSymbol("10", st.NORMAL, []float64{0, 0, 5, 20, 75}, []int{}))
	sm.Add(Slot.NewSymbol("9", st.NORMAL, []float64{0, 0, 5, 10, 50}, []int{}))
	sm.Add(Slot.NewSymbol("S", st.SCATTER, []float64{0, 0, 50, 0, 0}, []int{0, 0, 10, 15, 20}))
	sm.Add(Slot.NewSymbol("W", st.WILD, []float64{0, 0, 0, 0, 0}, []int{}))


	//新增LineGame計算 設置高度 滾輪 素材管理
	cl := Slot.NewCalcLine(3, rs, sm)
	cl.Init()

	result := cl.GetResult()
	fmt.Println("RTP:", result.GetRTP())
	fmt.Println("FreeP:", result.GetFreeP())
}
```

#### 架構圖
![架構](https://github.com/HyanSource/Shona/blob/master/ShonaArchitecture.png)

#### SymbolType `素材類型`
`NORMAL`:一般素材

`WILD`:百搭素材

`SCATTER`:免費素材

#### ISymbol `素材模塊`
```go
    //素材名稱 素材類型 賠率
    Slot.NewSymbol("K", st.NORMAL, []float64{0, 0, 5, 30, 100})
```

#### ISymbolManage `管理素材模塊`
```go
//新增素材類型以及賠率
    sm := NewSymbolManage()
    sm.Add(NewSymbol("K", st.NORMAL, []float64{0, 0, 5, 30, 100}))
```

#### Ireel `滾輪管理模塊`
```go
//新增5行滾輪條
	rs := []ISlot.Ireel{
		Slot.Newreel([]string{"K", "Q", "J", "10", "9", "9", "K", "Q", "J", "10", "9", "10", "Q", "J", "10", "9", "9", "J", "10", "9"}),
		Slot.Newreel([]string{"W", "K", "Q", "J", "10", "9", "S", "K", "Q", "J", "10", "9", "K", "Q", "J", "10", "9", "J", "J", "9"}),
		Slot.Newreel([]string{"W", "K", "Q", "J", "10", "9", "S", "K", "Q", "J", "10", "9", "K", "Q", "J", "10", "9", "J", "10", "9"}),
		Slot.Newreel([]string{"W", "K", "Q", "J", "10", "9", "S", "K", "Q", "J", "10", "9", "K", "Q", "J", "10", "9", "J", "10", "9"}),
		Slot.Newreel([]string{"W", "K", "Q", "J", "10", "9", "9", "K", "Q", "J", "10", "9", "K", "Q", "J", "10", "9", "J", "10", "9"}),
	}
```

#### ICalc `遊戲規則計算模塊`
```go
	//新增LineGame計算 設置高度 滾輪 素材管理
	cl := Slot.NewCalcLine(3, rs, sm)
	cl.Init()

```

#### IResult `數值模塊`
```go
	result := cl.GetResult()
	fmt.Println("RTP:", result.GetRTP())
	fmt.Println("FreeP:", result.GetFreeP())
```