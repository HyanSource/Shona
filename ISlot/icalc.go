package ISlot

/*規則計算*/
type ICalc interface {
	Init()
	GetLen() int
	GetHeight() int
	GetPossibility() int                        //總共機率
	GetGameTable(index int) ([]string, float64) //盤面
	GetRTP() float64                            //計算RTP
}
