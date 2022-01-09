package basic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// the value in map can be a function
func Test_map(t *testing.T) {
	m := map[int]func(n int)int{}
	m[1] = func(n int) int {
		return n
	}
	m[2] = func(n int) int {
		return n * n
	}
	m[3] = func(n int) int {
		return n * n * n
	}
	assert.Equal(t, 2, m[1](2))
	assert.Equal(t, 4, m[2](2))
	assert.Equal(t, 8, m[3](2))

}
