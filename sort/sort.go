// Package sort implements the following algorithms: selection sort, insertion sort, shell sort, merge sort.
// For merge sort, I don't know what kind of interface is appropriate, so it only works on int slices.
package sort

import (
	"fmt"
	"sort"
)

var _ = fmt.Println

// SelectionSort implements the selection algorithm.
func SelectionSort(data sort.Interface) {
	n := data.Len()

	for i := 0; i < n; i++ {
		minId := i
		// find the minium
		for j := i + 1; j < n; j++ {
			if data.Less(j, minId) {
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

func ShellSort(data sort.Interface) {
	n := data.Len()

	h := 1
	// 3x+1 increment sequence:  1, 4, 13, 40, 121, 364, 1093, ...
	for h < n/3 {
		h = 3*h + 1
	}

	for h >= 1 {
		// h-sort the array
		for i := h; i < n; i++ {
			for j := i; j >= h && data.Less(j, j-h); j -= h {
				data.Swap(j, j-h)
			}
		}
		h /= 3
	}
}

// merge merges 2 parts of the sorted slices using the auxiliary slice.
func merge(a, aux []int, lo, mid, hi int) {
	if a[mid] <= a[mid+1] { // no need to merge
		return
	}
	copy(aux[lo:], a[lo:mid+1])
	// reverse copy the other part, so don't need sentinel in merg
	// refer to Algorithms, 3ed.
	for i := mid + 1; i <= hi; i++ {
		aux[hi+mid+1-i] = a[i]
	}
	i := lo
	j := hi
	for k := lo; k <= hi; k++ {
		if aux[i] <= aux[j] {
			a[k] = aux[i]
			i++
		} else {
			a[k] = aux[j]
			j--
		}
	}
}

func mergeSort(a, aux []int, lo, hi int) {
	if hi <= lo {
		return
	}
	mid := (hi-lo)/2 + lo
	mergeSort(a, aux, lo, mid)
	mergeSort(a, aux, mid+1, hi)
	merge(a, aux, lo, mid, hi)
}

func MergeSort(a []int) {
	aux := make([]int, len(a))
	mergeSort(a, aux, 0, len(a)-1)
}
