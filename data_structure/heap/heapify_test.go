package heap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxHeapify(t *testing.T) {
	nums := []int{4}
	MaxHeapify(nums, 0)
	assert.Equal(t, 4, nums[0])

	nums = append(nums, 10)
	MaxHeapify(nums, 0)
	assert.Equal(t, 10, nums[0])

	nums = append(nums, 16)
	MaxHeapify(nums, 0)
	assert.Equal(t, 16, nums[0])

	nums = append(nums, 14, 7, 9, 3, 2, 8, 1)
	MaxHeapify(nums, 1)
	assert.Equal(t, 14, nums[1])
	assert.Equal(t, 4, nums[8])
}
