package basic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_defer(t *testing.T) {
	f := func(flag *int) {
		// defer stack
		i := 0
		defer func() {
			*flag = 2
			t.Logf("defer1: %d, i:%d", *flag, i)
		}()
		defer func() {
			recover()
			*flag = 3
			t.Logf("defer2: %d, i:%d", *flag, i)
		}()
		*flag = 1
		i++
		panic("panic panic")
		return
	}
	flag := 0
	f(&flag)
	assert.Equal(t, 2, flag)
}
