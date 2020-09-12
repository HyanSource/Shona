package ISlot

/*規則計算*/
type ICalc interface {
	Init()
	GetLen() int
	GetPossibility() int             //總共機率
	GetGameTable(index int) []string //盤面

	CalcOdds(result []string) float64                //計算賠率
	CalcScatter(result []string) (float64, int, int) //計算免費賠率和次數

	GetResult() IResult //計算RTP

	GetNormalPlayResult(playcount int) IResult //模擬正常玩的RTP
}
