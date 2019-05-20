package stringtranspose

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranspose(t *testing.T) {
	data := [][]string{
		[]string{"a1", "a2", "a3"},
		[]string{"b1", "b2", "b3"},
		[]string{"c1", "c2", "c3"},
	}
	result := Transpose(data)
	assert.Len(t, result, 3)
	assert.Len(t, result[0], 3)
	assert.Len(t, result[1], 3)
	assert.Len(t, result[2], 3)
	assert.Exactly(t, []string{"a1", "b1", "c1"}, result[0])
	assert.Exactly(t, []string{"a2", "b2", "c2"}, result[1])
	assert.Exactly(t, []string{"a3", "b3", "c3"}, result[2])
}

func TestTransposeIrregular(t *testing.T) {
	data := [][]string{
		[]string{"a1", "a2", "a3"},
		[]string{"b1", "b2"},
		[]string{"c1", "c2", "c3", "c4"},
	}
	result := TransposeIrregular(data)
	assert.Len(t, result, 4)
	assert.Len(t, result[0], 3)
	assert.Len(t, result[1], 3)
	assert.Len(t, result[2], 3)
	assert.Exactly(t, []string{"a1", "b1", "c1"}, result[0])
	assert.Exactly(t, []string{"a2", "b2", "c2"}, result[1])
	assert.Exactly(t, []string{"a3", "", "c3"}, result[2])
	assert.Exactly(t, []string{"", "", "c4"}, result[3])
}
