package basic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_array(t *testing.T) {
	a1 := [...]int{1 ,2, 3}
	a2 := [...]int{1 ,2, 3}
	assert.Equal(t, true, a1 == a2)

	a3 := [...]int{1, 2, 4}
	assert.Equal(t, false, a1 == a3)

	//a4 := [...]int{1, 2}
	//assert.Equal(t, false, a1 == a4)
}
