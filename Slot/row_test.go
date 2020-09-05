package Slot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRow(t *testing.T) {
	r := NewRow([]string{"A", "A", "K", "Q", "J", "10"})

	assert.Equal(t, 6, r.GetLen())
	assert.Equal(t, []string{"A", "A", "K", "Q", "J", "10"}, r.GetRow())
	assert.Equal(t, []string{"A", "K", "Q"}, r.GetSymbols(1, 3))

}
