package sort

import (
	"github.com/cyfdecyf/goutil/rand"
	"sort"
	"testing"
)

var data [][]int

func init() {
	const nSeq = 10
	data = make([][]int, nSeq, nSeq)

	for i := 0; i < nSeq; i++ {
		n := int(1 << uint32(i))
		data[i] = rand.GenKRandomLessN(n, n)
	}
}

// copyTestData creates a copy of the test data. Use this to ensure that each
// sort function is working on the same data, and its useful to print the data
// before sorting.
func copyTestData() [][]int {
	nSeq := len(data)
	cp := make([][]int, nSeq, nSeq)
	for i := 0; i < nSeq; i++ {
		n := len(cp[i])
		cp[i] = make([]int, n, n)
		copy(cp[i], data[i])
	}
	return cp
}

func TestSelectionSort(t *testing.T) {
	for i, d := range copyTestData() {
		is := sort.IntSlice(d)
		SelectionSort(is)
		if !sort.IsSorted(is) {
			t.Fatalf("SelectionSort wrong for data: %v, get: %v", data[i], d)
		}
	}
}
