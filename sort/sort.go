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

func InsertionSort(data sort.Interface) {
	n := data.Len()
	for i := 1; i < n; i++ {
		for j := i; j > 0 && data.Less(j, j-1); j-- {
			data.Swap(j, j-1)
		}
	}
}
