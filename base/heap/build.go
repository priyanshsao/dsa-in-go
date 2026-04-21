package heap

import (
	"container/heap"

	"github.com/sirupsen/logrus"
)

// fromInit uses heap.Init to build array, time complexity: O(n)
func Init(array []int, ht HeapType) *Heap {
	h := new(Heap)

	// cannot assign array directly
	// if h.data is modified then
	// so does the variable array
	h.data = append([]int{}, array...)
	h.typeOf = ht

	// initial rearrangement is done
	// the slice is rearranged according to
	// heap arrangement
	heap.Init(h)
	logrus.Debug("heap initialized: ", h.Data())

	return h
}

// time complexity: O(nlogn), which is greater than O(n)
func InitByPush(array []int, ht HeapType) *Heap {
	h := new(Heap)
	h.typeOf = ht

	for _, v := range array {
		heap.Push(h, v)
	}
	logrus.Debug("heap initialized by push: ", h.Data())

	return h
}