package ISlot

/*計算產生的結果*/
type IResult interface {
	GetRTP() float64   //RTP
	GetFreeP() float64 //免費觸發機率
}
