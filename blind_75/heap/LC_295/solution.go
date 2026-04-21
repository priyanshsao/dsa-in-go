package main

import (
	stdheap "container/heap"

	"github.com/priyanshsao/dsa-in-go/base/heap"
)

type MedianFinder struct {
    smHeap *heap.Heap
	lgHeap *heap.Heap
}


func Constructor() MedianFinder {
    mf := new(MedianFinder)
	mf.smHeap = heap.Init([]int{}, heap.Max_Heap)
	mf.lgHeap = heap.Init([]int{}, heap.Min_Heap)

	return *mf
}


func (mf *MedianFinder) AddNum(num int)  {
	stdheap.Push(mf.smHeap, num)

	stdheap.Push(mf.lgHeap, stdheap.Pop(mf.smHeap))

	if len(mf.lgHeap.Data()) > len(mf.smHeap.Data()) {
		stdheap.Push(mf.smHeap, stdheap.Pop(mf.lgHeap))
	}
}


func (mf *MedianFinder) FindMedian() float64 {
    if len(mf.smHeap.Data()) > len(mf.lgHeap.Data()) {
		return float64(mf.smHeap.Data()[0])
	}

	return float64((mf.smHeap.Data()[0] + mf.lgHeap.Data()[0]))/ 2.0
}
