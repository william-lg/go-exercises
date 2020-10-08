package heap

func Sort(nums []int) {
	BuildMaxHeap(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		temp := nums[:i]
		MaxHeapify(temp, 0)
	}
}
