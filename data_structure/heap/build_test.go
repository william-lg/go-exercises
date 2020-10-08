package heap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildMaxHeap(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	BuildMaxHeap(nums)
	assert.Equal(t, 5, nums[0])
	assert.Equal(t, 1, nums[3])
	assert.Equal(t, 2, nums[4])
}
