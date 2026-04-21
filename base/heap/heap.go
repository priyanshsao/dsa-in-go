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

package heap

type HeapType int

const (
	Min_Heap HeapType = iota
	Max_Heap
)

type Heap struct {
	data []int
	typeOf HeapType
}

// returns the length of heap(indirectly: length of array)
func (h Heap) Len() int {
	return len(h.data)
}

func (h Heap) Less(i, j int) bool {
	if h.typeOf == Max_Heap {
		return h.data[j] < h.data[i]
	}
	return h.data[i] < h.data[j]
}

func (h Heap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *Heap) Push(v any) {
	h.data = append(h.data, v.(int))
}

func (h *Heap) Pop() any {
	old := h.data
	l := len(old)
	v := old[l-1]
	h.data = old[:l-1]
	return v
}

func (h *Heap) Modify(i, v int) {
	h.data[i] = v
}

func (h *Heap) Data() []int {
	return h.data
}