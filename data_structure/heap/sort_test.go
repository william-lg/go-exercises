package heap

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	nums := []int{9, 8, 3, 4, 5, 3, 3, 6, 4, 7, 2, 1}
	Sort(nums)
	for _, v := range nums {
		fmt.Printf("%v ", v)
	}
}
