// Package sort implements the following algorithms: selection sort, insertion
// sort, shell sort, merge sort, quicksort. For merge sort, I don't know what kind of
// interface is appropriate, so it only works on int slices.
package sort

import (
	"fmt"
	"github.com/cyfdecyf/goutil/rand"
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

// MergeSort does merge sort on slice a.
func MergeSort(a []int) {
	aux := make([]int, len(a))
	mergeSort(a, aux, 0, len(a)-1)
}

// merge merges 2 parts of the sorted slices into the auxiliary slice.
func mergeInto(a, aux []int, lo, mid, hi int) {
	// fmt.Printf("Merging a: %v into %v\n", a, aux)
	if a[mid] <= a[mid+1] { // no need to merge
		copy(aux[lo:hi+1], a[lo:hi+1])
		return
	}
	i := lo
	j := mid + 1
	for k := lo; k <= hi; k++ {
		if i > mid {
			copy(aux[k:], a[j:hi+1])
			break
		} else if j > hi {
			copy(aux[k:], a[i:mid+1])
			break
		}
		if a[i] <= a[j] {
			aux[k] = a[i]
			i++
		} else {
			aux[k] = a[j]
			j++
		}
	}
}

// mergeSortNoCopy sorts slice a, puts the results in aux.
func mergeSortNoCopy(a, aux []int, lo, hi int) {
	// fmt.Printf("merge sort %v into %v, [%d, %d]\n", a, aux, lo, hi)
	if hi <= lo {
		return
	}
	mid := (hi-lo)/2 + lo
	// For merge sort, all the data movement happens in merge. If we switch
	// the role of a and aux for each merge, then we can avoid the copy.
	mergeSortNoCopy(aux, a, lo, mid)
	mergeSortNoCopy(aux, a, mid+1, hi)
	mergeInto(a, aux, lo, mid, hi)
}

// MergeSortNoCopy does merge sort on slice a, without copying for each merge.
func MergeSortNoCopy(a []int) {
	aux := make([]int, len(a))
	copy(aux, a)
	mergeSortNoCopy(aux, a, 0, len(a)-1)
}

// MergeSortBU does bottom up merge sort on slice a.
func MergeSortBU(a []int) {
	n := len(a)
	aux := make([]int, n)
	// sz is the size of each sorted part
	for sz := 1; sz < n; sz += sz {
		// lo is the start index of the first sorted part
		for lo := 0; lo < n-sz; lo += sz + sz {
			hi := lo + sz + sz - 1
			if hi > n-1 {
				hi = n - 1
			}
			merge(a, aux, lo, lo+sz-1, hi)
		}
	}
}

// Partition the subslice a[lo:hi+1] by returning an index j
// so that a[lo:j] <= a[j] <= a[j+1:hi+1].
func Partition(data sort.Interface, lo, hi int) int {
	i := lo + 1
	j := hi
	pivot := lo

	for {
		// Problem of this partition:  i, j stops at item which is the same as
		// pivot.
		for data.Less(i, pivot) && i < hi { // data[i] < data[pivot]
			i++
		}
		for data.Less(pivot, j) && j > lo { // data[pivot] < data[j]
			j--
		}
		if i >= j {
			break
		}
		data.Swap(i, j)
	}
	// At this point, we must have data[j] <= data[pivot], that's why we swap j and pivot.
	data.Swap(pivot, j)
	return j
}

func quickSort(data sort.Interface, lo, hi int) {
	if hi <= lo {
		return
	}
	j := Partition(data, lo, hi)
	quickSort(data, lo, j-1)
	quickSort(data, j+1, hi)
}

func QuickSort(data sort.Interface) {
	rand.Shuffle(data)
	quickSort(data, 0, data.Len()-1)
}
