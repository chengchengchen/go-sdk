package basic

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func Test_string(t *testing.T) {
	s := "我"
	assert.Equal(t, 3, len(s))

	s2, err := strconv.Atoi("123")
	assert.Equal(t, nil, err)
	assert.Equal(t, 123, s2)

	s3, err2 := strconv.Atoi("123hj")
	assert.NotEqual(t, nil, err2)
	assert.Equal(t, 0, s3) // default value in int
}
