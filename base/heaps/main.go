package main

import (
	"container/heap"

	"github.com/sirupsen/logrus"
)

func main() {

	// remove unwanted things and enforce colors
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true,
		ForceColors:      true,
		PadLevelText:     true,
	})

	// this ensures all the logs are visible in stdout
	logrus.SetLevel(logrus.DebugLevel)

	//build array
	initS := []int{2, 74, 0, 89, 14, 35}
	// pushS := []int{2, 74, 0, 89, 14, 35}

	// use init function
	iHeap := FromInit(initS)

	// apply pop method to return min element
	min := heap.Pop(iHeap)
	logrus.Warn("element poped: ", min)
	logrus.Debug("heap after pop: ", *iHeap)

	(*iHeap)[2] = 0
	logrus.Info("change in value heap[2]:0")
	logrus.Debug("heap after modify: ", *iHeap)

	heap.Fix(iHeap, 2)
	logrus.Debug("heap after fix: ", *iHeap)
	// use push function
	// pHeap := fromPush(pushS)

	logrus.Warn("element index:0 removed")
	heap.Remove(iHeap, 0)
	logrus.Debug("heap after removal: ", *iHeap)
}

// takes O(n) time
func FromInit(array []int) *MinHeap {
	//store the array in heap
	h := new(MinHeap)
	*h = array

	// initial rearrangement is done
	// the slice is rearranged according to 
	// heap arrangement
	heap.Init(h)
	logrus.Debug("init heap called: ", *h)

	return h
}

func fromPush(array []int) *MinHeap {
	
	h := new(MinHeap)

	for _, v := range array {
		heap.Push(h, v)
	}

	logrus.Debug("heap created by push: ", *h)

	return h
}