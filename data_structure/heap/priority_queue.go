package heap

import "math"

type MaxPriorityQueue interface {
	Insert(key int)
	Maximum() int
	ExtractMax() int
	IncreaseKey(elem, key int)
}

type ArrayMaxPQ []int

func (a *ArrayMaxPQ) Insert(key int) {
	*a = append(*a, math.MinInt32)
	a.IncreaseKey(len(*a)-1, key)
}

func (a *ArrayMaxPQ) Maximum() int {
	return (*a)[0]
}

func (a *ArrayMaxPQ) ExtractMax() int {
	if len(*a) == 0 {
		panic("priority queue is empty")
	}

	ti := len(*a) - 1
	(*a)[0], (*a)[ti] = (*a)[ti], (*a)[0]
	tv := (*a)[ti]
	*a = (*a)[:ti]
	MaxHeapify(*a, 0)

	return tv
}

func (a *ArrayMaxPQ) IncreaseKey(elem, key int) {
	(*a)[elem] = key
	for i := elem; i > 0; i = Parent(i) {
		MaxHeapify(*a, i)
	}
	MaxHeapify(*a, 0)
}
