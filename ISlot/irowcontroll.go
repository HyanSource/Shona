package ISlot

type IRowsControll interface {
	GetRowsLen() int    //取得滾輪長度
	GetAllRows() []IRow //取得所有滾輪

	GetPossibility() int //取得所有機率發生次數

	GetGameTable(index int) ([]string, float64) //取得特定盤面

	GetLineGameRTP() float64 //取得RTP
	GetWayGameRTP() float64  //取得RTP
}
