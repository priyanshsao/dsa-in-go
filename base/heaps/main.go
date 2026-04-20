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

	//slice for building heap
	initSlc := []int{2, 74, 0, 89, 14, 35}

	// use init function.
	// iminHeap = min heap created using fromInit()
	iMinHeap := InitMinHeap(initSlc)

	// could also use push function
	// pHeap := fromPush(pushS)
	// but init is more efficient

	// apply pop method to return min element
	min := heap.Pop(iMinHeap)
	logrus.Warn("element poped: ", min)
	logrus.Debug("heap after pop: ", iMinHeap)

	(*iMinHeap)[2] = 0
	logrus.Info("change in value heap[2]=0")
	logrus.Debug("heap after modify: ", *iMinHeap)

	heap.Fix(iMinHeap, 2)
	logrus.Debug("heap after fix: ", *iMinHeap)

	logrus.Warn("element index:0 removed")
	heap.Remove(iMinHeap, 0)
	logrus.Debug("heap after removal: ", *iMinHeap)

	iMaxHeap := InitMaxHeap(initSlc)

	max := heap.Pop(iMaxHeap)
	logrus.Warn("element poped: ", max)
	logrus.Debug("heap after pop: ", *iMaxHeap)

	(*iMaxHeap)[2] = 99
	logrus.Info("change in value heap[2]=99")
	logrus.Debug("heap after modify: ", *iMaxHeap)

	heap.Fix(iMaxHeap, 2)
	logrus.Debug("heap after fix: ", *iMaxHeap)

	logrus.Warn("element index=0 removed")
	heap.Remove(iMaxHeap, 0)
	logrus.Debug("heap after removal: ", *iMaxHeap)

}
