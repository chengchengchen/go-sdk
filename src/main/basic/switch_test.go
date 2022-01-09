package basic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_switch(t *testing.T) {
	a := 9
	res := 0
	switch {
	case a <= 8:
		res += 1
	case a >= 6:
		res += 2
	case a <= 9:
		res += 3
	default:
		res += 4
	}
	assert.Equal(t, 2, res)


	// test the function of fallthrough
	res = 0
	switch {
	case a <= 8:
		res += 1
		fallthrough
	case a >= 6:
		res += 2
		fallthrough
	case a <= 9:
		res += 3
		fallthrough
	default:
		res += 4
	}
	assert.Equal(t, 9, res)
}
