package heap

func BuildMaxHeap(nums []int) {
	lth := len(nums)
	for i := Parent(lth - 1); i >= 0; i-- {
		MaxHeapify(nums, i)
	}
}
