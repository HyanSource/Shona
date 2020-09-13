package ISlot

import (
	"github.com/HyanSource/Shona/st"
)

type ISymbol interface {
	SetSymbol()
	GetName() string
	GetSymbolType() string
	GetOdds() []float64
	GetFree() []int
	CheckType(st st.ST) bool
}
