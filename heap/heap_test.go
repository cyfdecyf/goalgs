package heap

import (
	"github.com/cyfdecyf/goutil/rand"
	"testing"
)

var data [][]int

func init() {
	const nSeq = 12
	data = make([][]int, nSeq)

	data[0] = []int{0}
	data[1] = []int{1, 0}

	for i := 2; i < nSeq; i++ {
		n := int(1<<uint32(i)) + rand.Rand(0, 5)
		data[i] = rand.GenKRandomLessN(n, n)
	}
}

func TestHeap(t *testing.T) {
	for _, d := range data {
		h := NewHeap(CompareIntReverse)
		for _, v := range d {
			h.Push(v)
		}
		for i := 0; i < len(d); i++ {
			v := h.Pop().(int)
			if v != i {
				t.Fatalf("Heap sort error for data %v at %dth pop", d, i)
			}
		}
	}
}

func TestHeapFromSlice(t *testing.T) {
	for _, d := range data {
		s := make([]interface{}, len(d))
		for i := 0; i < len(d); i++ {
			s[i] = d[i]
		}
		h := NewHeapFromSlice(CompareIntReverse, s)
		for i := 0; i < len(d); i++ {
			v := h.Pop().(int)
			if v != i {
				t.Fatalf("Heap sort with NewHeapFromSlice error for data %v at %dth pop", d, i)
			}
		}
	}
}
