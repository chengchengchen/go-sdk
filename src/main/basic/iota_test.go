package basic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_iota(t *testing.T) {
	assert.Equal(t, 1, i)
	assert.Equal(t, 2, j)
	assert.Equal(t, 3, k)

	assert.Equal(t, 1, one)
	assert.Equal(t, 2, two)
	assert.Equal(t, 4, three)
	assert.Equal(t, "four", four)
	assert.Equal(t, "four", five)
	assert.Equal(t, 32, six)
}
