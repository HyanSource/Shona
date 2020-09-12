package Slot

import "github.com/HyanSource/Shona/ISlot"

func NewResult(count int, bet float64, winmoney float64, freecount int) ISlot.IResult {
	return &Result{
		count:     count,
		bet:       bet,
		winmoney:  winmoney,
		freecount: freecount,
	}
}

type Result struct {
	count     int     //遊玩次數
	bet       float64 //每次下注金額
	winmoney  float64 //總共獲得金額
	freecount int     //總共獲得免費的次數
}

//RTP
func (t *Result) GetRTP() float64 {
	return t.winmoney / (float64(t.count) * t.bet)
}

//免費觸發機率
func (t *Result) GetFreeP() float64 {
	return float64(t.freecount) / float64(t.count)
}
