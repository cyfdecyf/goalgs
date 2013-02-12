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
	const nSeq = 10
	data = make([][]int, nSeq, nSeq)

	for i := 0; i < nSeq; i++ {
		n := int(1<<uint32(i)) + rand.Rand(0, 10)
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
