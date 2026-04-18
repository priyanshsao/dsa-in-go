// To build heap

// following functions should be satisfied by your type

// <type>.Init() 
// work: convert an array to heap
// time complexity:  O(n)

// <type>.Push()
// work: Push a new element onto the heap and call Fix() to reorder
// time complexity:  O(logn)

// <type>.Pop()
// work: remove and return an element followed by updating the order
// time complexity:  O(logn)

// <type>.Fix()
// work: used for updating the order according to heap type(min/max)
// time complexity:  O(logn)

// <type>.Remove()
// work: remove and return an element at index i from the heap and update heap
// time complexity:  O(logn)

package main

type MinHeap []int

// returns the length of heap(indirectly: array)
func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(v any) {
	*h = append(*h, v.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	l := len(old)
	v := old[l-1]
	*h = old[:l-1]
	return v
}