package ISlot

type IRow interface {
	SetSymbol(s []string)
	GetLen() int
	GetRow() []string
	GetSymbols(index int, count int) []string
}
