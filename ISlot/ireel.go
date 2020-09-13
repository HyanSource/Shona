package ISlot

type Ireel interface {
	SetSymbol(s []string)
	GetLen() int
	Getreel() []string
	GetSymbols(index int, count int) []string
}
