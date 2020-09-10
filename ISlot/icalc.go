package ISlot

type ICalc interface {
	GetGameTable() []string
	GetGameOdds() float64
	GetRTP() float64 //計算RTP
}
