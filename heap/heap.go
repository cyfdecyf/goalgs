// Package heap provides heap implementation. The heap in this package uses
// interface{} to hold objects and requires a compare function.
package heap

import (
	"fmt"
)

var _ = fmt.Println

// Comparator compares two objects, should return <0, 0, >0 for less than, equal and
// larger than.
type Comparator func(interface{}, interface{}) int

// Heap provides a binary heap. The top most element is the largest one.
type Heap struct {
	comp func(interface{}, interface{}) int
	item []interface{} // indicies start at 1
}

func CompareInt(i interface{}, j interface{}) int {
	ii := i.(int)
	jj := j.(int)
	return ii - jj
}

func CompareIntReverse(i interface{}, j interface{}) int {
	ii := i.(int)
	jj := j.(int)
	return jj - ii
}

// NewPQ returns a priority queue with given comparator.
func NewHeap(cmp Comparator) *Heap {
	return &Heap{comp: cmp, item: make([]interface{}, 1)}
}

func NewHeapFromSlice(cmp Comparator, data []interface{}) *Heap {
	h := &Heap{comp: cmp, item: make([]interface{}, 1+len(data))}
	copy(h.item[1:], data)
	n := len(h.item) - 1
	for k := n / 2; k >= 1; k-- {
		h.sink(k)
	}
	return h
}

// Insert one object.
func (h *Heap) Push(v interface{}) {
	h.item = append(h.item, v)
	h.swim(len(h.item) - 1)
}

// Pop removes the largest object from the heap.
func (h *Heap) Pop() interface{} {
	if len(h.item) == 1 {
		return nil
	}
	v := h.item[1]
	h.item[1] = h.item[len(h.item)-1]
	h.item = h.item[:len(h.item)-1]
	h.sink(1)
	return v
}

// Move item k up to the correct place.
func (h *Heap) swim(k int) {
	for k > 1 && h.comp(h.item[k/2], h.item[k]) < 0 {
		h.item[k/2], h.item[k] = h.item[k], h.item[k/2]
		k /= 2
	}
}

// Move item k down to the correct place.
func (h *Heap) sink(k int) {
	n := len(h.item) - 1
	for 2*k <= n {
		j := 2 * k
		// exchange with larger child
		if j < n && h.comp(h.item[j], h.item[j+1]) < 0 {
			j++
		}
		if h.comp(h.item[k], h.item[j]) >= 0 {
			break
		}
		h.item[k], h.item[j] = h.item[j], h.item[k]
		k = j
	}
}
