package Slot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLineManage(t *testing.T) {
	lm := NewLineManage(5, 3)

	assert.True(t, lm.SetLine([]int{0, 1, 2, 3, 4}))
	assert.True(t, lm.SetLine([]int{5, 6, 7, 8, 9}))
	assert.True(t, lm.SetLine([]int{10, 11, 12, 13, 14}))

	assert.Equal(t, 3, lm.Len())

	assert.Equal(t, []int{0, 1, 2, 3, 4}, lm.GetLine(0))
	assert.Equal(t, []int{5, 6, 7, 8, 9}, lm.GetLine(1))
	assert.Equal(t, []int{10, 11, 12, 13, 14}, lm.GetLine(2))
}
