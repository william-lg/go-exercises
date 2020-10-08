package heap

func MaxHeapify(nums []int, idx int) {
	var (
		lth       = len(nums)
		leftIdx   = Left(idx)
		rightIdx  = Right(idx)
		largerIdx = idx
	)

	// 已经是叶子节点
	if leftIdx >= lth {
		return
	}
	if nums[leftIdx] > nums[largerIdx] {
		largerIdx = leftIdx
	}
	if rightIdx < lth && nums[rightIdx] > nums[largerIdx] {
		largerIdx = rightIdx
	}

	if largerIdx != idx {
		nums[idx], nums[largerIdx] = nums[largerIdx], nums[idx]
		MaxHeapify(nums, largerIdx)
	}
}
