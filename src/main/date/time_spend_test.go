package date

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_timeSpend(t *testing.T) {
	v := func(res *int) {
		for i := 0; i < 100000; i++ {
			*res += i
		}
	}
	assert.Equal(t, 4999950000, timeSpend(v))
}
