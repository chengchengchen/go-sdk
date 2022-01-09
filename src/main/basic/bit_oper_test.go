package basic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_byte_operation(t *testing.T) {
	a := 2 // 0010
	b := 3 // 0011
	assert.Equal(t, 0, a &^ b)
	assert.Equal(t, 1, b &^ a)
}