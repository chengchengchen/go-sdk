package basic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_slice(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} // initialize directly
	assert.Equal(t, 9, len(nums))
	assert.Equal(t, 9, cap(nums))
	//num1 := nums[2:3]

	nums2 := make([]int, 0)
	assert.Equal(t, 0, cap(nums2))
	for i := 1; i < 10; i++ {
		nums2 = append(nums2, i)
	}
	assert.Equal(t, 9, len(nums2))
	assert.Equal(t, 16, cap(nums2))

	nums3 := nums2[2:]
	for i := 10; i < 20; i++ {
		nums2 = append(nums2, i) // nums2 resize and point to another memory after appending 16
	}
	assert.Equal(t, 7, len(nums3))
	assert.Equal(t, 14, cap(nums3))
	// modify the value in nums3, it will not modify the value in nums2. Because the slices of nums2 and nums3 are different.
	nums3[0] = 100
	assert.Equal(t, 3, nums2[2])
	t.Logf("%v\n", nums)
	t.Logf("%v\n", nums2)
	t.Logf("%v\n", nums3)
}
