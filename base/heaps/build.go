package main

import (
	"container/heap"

	"github.com/sirupsen/logrus"
)

// fromInit uses heap.Init to build array, time complexity: O(n)
func InitMinHeap(array []int) *MinHeap {
	h := new(MinHeap)
	*h = array

	// initial rearrangement is done
	// the slice is rearranged according to
	// heap arrangement
	heap.Init(h)
	logrus.Debug("init heap called: ", *h)

	return h
}

// time complexity: O(nlogn), which is greater than O(n)
func InitMinHeapFromPush(array []int) *MinHeap {
	h := new(MinHeap)

	for _, v := range array {
		heap.Push(h, v)
	}
	logrus.Debug("heap created by push: ", h)

	return h
}

// fromInit uses heap.Init to build array, time complexity: O(n)
func InitMaxHeap(array []int) *MaxHeap {
	h := new(MaxHeap)
	*h = array

	// initial rearrangement is done
	// the slice is rearranged according to
	// heap arrangement
	heap.Init(h)
	logrus.Debug("init heap called: ", *h)

	return h
}

// time complexity: O(nlogn), which is greater than O(n)
func InitMaxHeapFromPush(array []int) *MaxHeap {
	h := new(MaxHeap)

	for _, v := range array {
		heap.Push(h, v)
	}
	logrus.Debug("heap created by push: ", h)

	return h
}
