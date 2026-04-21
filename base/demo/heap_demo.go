// This main file is for understanding heap.
// The other files are used to create heap,
// and perform operations in other projects
package main

import (
	stdheap "container/heap"

	"github.com/priyanshsao/dsa-in-go/base/heap"
	"github.com/sirupsen/logrus"
)

func main() {

	formatLogs()

	//slice for building heap
	initSlc := []int{2, 74, 0, 89, 14, 35}

	// use init function.
	// iminHeap = min heap created using fromInit()
	iMinHeap := heap.Init(initSlc, heap.Min_Heap)

	// could also use push function
	// pHeap := InitByPush(pushS, Min_Heap)
	// but init is more efficient

	// apply pop method to return min element
	min := stdheap.Pop(iMinHeap)
	logrus.Warn("element poped: ", min)
	logrus.Debug("heap after pop: ", iMinHeap.Data())

	iMinHeap.Modify(2, 0)
	logrus.Info("change in value heap[2]=0")
	logrus.Debug("heap after modify: ", iMinHeap.Data())

	stdheap.Fix(iMinHeap, 2)
	logrus.Debug("heap after fix: ", iMinHeap.Data())

	logrus.Warn("element index:0 removed")
	stdheap.Remove(iMinHeap, 0)
	logrus.Debug("heap after removal: ", iMinHeap.Data())

	iMaxHeap := heap.Init(initSlc, heap.Max_Heap)

	max := stdheap.Pop(iMaxHeap)
	logrus.Warn("element poped: ", max)
	logrus.Debug("heap after pop: ", iMaxHeap.Data())

	iMaxHeap.Modify(2, 99)
	logrus.Info("change in value heap[2]=99")
	logrus.Debug("heap after modify: ", iMaxHeap.Data())

	stdheap.Fix(iMaxHeap, 2)
	logrus.Debug("heap after fix: ", iMaxHeap.Data())

	logrus.Warn("element index=0 removed")
	stdheap.Remove(iMaxHeap, 0)
	logrus.Debug("heap after removal: ", iMaxHeap.Data())

}

func formatLogs() {
	
	// remove unwanted things and enforce colors
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true,
		ForceColors:      true,
		PadLevelText:     true,
	})

	// this ensures all the logs are visible in stdout
	logrus.SetLevel(logrus.DebugLevel)
}