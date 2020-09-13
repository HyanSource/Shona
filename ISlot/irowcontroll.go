package ISlot

type IreelsControll interface {
	GetreelsLen() int    //取得滾輪長度
	GetAllreels() []Ireel //取得所有滾輪

	GetPossibility() int //取得所有機率發生次數

	GetGameTable(index int) ([]string, float64) //取得特定盤面
}
