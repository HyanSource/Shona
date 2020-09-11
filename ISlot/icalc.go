package ISlot

/*規則計算*/
type ICalc interface {
	Init()
	GetLen() int
	GetPossibility() int             //總共機率
	GetGameTable(index int) []string //盤面

	CalcOdds(result []string) float64           //計算賠率
	CalcScatter(result []string) (float64, int) //計算免費賠率和次數

	GetRTP() float64   //計算RTP
	GetFreeP() float64 //取得免費觸發機率

	GetRandomPlayRTP(playcount int) float64 //模擬正常玩的RTP
}
