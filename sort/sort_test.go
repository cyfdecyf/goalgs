package sort

import (
	"fmt"
	"github.com/cyfdecyf/goutil/rand"
	"sort"
	"testing"
)

var _ = fmt.Println

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

// copyTestData creates a copy of the test data. Use this to ensure that each
// sort function is working on the same data, and its useful to print the data
// before sorting.
func copyTestData() [][]int {
	nSeq := len(data)
	cp := make([][]int, nSeq)
	for i := 0; i < nSeq; i++ {
		n := len(data[i])
		cp[i] = make([]int, n, n)
		copy(cp[i], data[i])
	}
	return cp
}

func testSortFunc(sortFunc func(sort.Interface), funcName string, t *testing.T) {
	for i, d := range copyTestData() {
		is := sort.IntSlice(d)
		sortFunc(is)
		for di, dv := range d {
			if di != dv {
				t.Fatalf("%s wrong for data: %v, get: %v", funcName, data[i], d)
			}
		}
	}
}

func TestSelectionSort(t *testing.T) {
	testSortFunc(SelectionSort, "SelectionSort", t)
}

func TestInsertionSort(t *testing.T) {
	testSortFunc(InsertionSort, "InsertionSort", t)
}

func TestShellSort(t *testing.T) {
	testSortFunc(ShellSort, "ShellSort", t)
}

func testMerge(funcName string, mf func(a, aux []int, lo, mid, hi int), t *testing.T) {
	testData := []struct {
		orig   []int
		merged []int
		lo     int
		mid    int
		hi     int
	}{
		{[]int{1, 0}, []int{0, 1}, 0, 0, 1},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 0, 5, 9},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 0, 4, 9},
		{[]int{4, 5, 6, 7, 8, 9, 0, 1, 2, 3}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 0, 5, 9},
		{[]int{4, 5, 6, 7, 8, 9, 0, 1, 2, 3}, []int{4, 5, 6, 0, 1, 7, 8, 9, 2, 3}, 3, 5, 7},
		{[]int{9, 0, 1, 2, 3, 4, 5, 6, 7, 8}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 0, 0, 9},
	}

	aux := make([]int, 10)
	for _, td := range testData {
		cp := make([]int, 10)
		copy(cp, td.orig)
		mf(cp, aux, td.lo, td.mid, td.hi)
		for i := td.lo; i <= td.hi; i++ {
			if cp[i] != td.merged[i] {
				t.Fatalf("%s error for data %v lo=%d mid=%d hi=%d, got %v, should be %v\n",
					funcName, td.orig, td.lo, td.mid, td.hi, cp, td.merged)
			}
		}
	}
}

func TestMerge(t *testing.T) {
	testMerge("merge", merge, t)
}

func TestMergeInto(t *testing.T) {
	mf := func(a, aux []int, lo, mid, hi int) {
		mergeInto(a, aux, lo, mid, hi)
		// copy aux back to a, so the test code can check correctness
		copy(a[lo:hi+1], aux[lo:hi+1])
	}
	testMerge("mergeInto", mf, t)
}

func testMergeSort(funcName string, ms func([]int), t *testing.T) {
	for i, d := range copyTestData() {
		ms(d)
		for di, dv := range d {
			if di != dv {
				t.Fatalf("%s wrong for data[%d]: %v, get: %v", funcName, i, data[i], d)
			}
		}
	}

}

func TestMergeSort(t *testing.T) {
	testMergeSort("MergeSort", MergeSort, t)
}

func TestMergeSortNoCopy(t *testing.T) {
	testMergeSort("MergeSortNoCopy", MergeSortNoCopy, t)
}

func TestMergeSortBU(t *testing.T) {
	testMergeSort("MergeSortBU", MergeSortBU, t)
}
