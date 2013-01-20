// Package sort implements the following algorithms: insertion sort
package sort

import (
	"sort"
)

// SelectionSort implements the selection algorithm.
func SelectionSort(data sort.Interface) {
	n := data.Len()

	for i := 0; i < n; i++ {
		minId := i
		// find the minium
		for j := i + 1; j < n; j++ {
			if data.Less(minId, j) {
				minId = j
			}
		}
		data.Swap(i, minId)
	}
}
