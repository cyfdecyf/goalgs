// Package uf implements 3 union-find algorithms: quick-find, quick-union and
// weighted quick-union.
package uf

import (
// "fmt"
)

// Union-find API
type UF interface {
	// Union add connection between p and q.
	Union(p, q int)
	// Find return component identifier for p.
	Find(p int) int
	// Connected return true if p and q are in the same component.
	Connected(p, q int) bool
	// Count return number of components.
	Count() int
}

// QuickFind implements the quick-find algorithm.
type QuickFind struct {
	id    []int // p and q are connected iff they have the same id in the array
	count int   // number of components
}

// NewQuickFind creates a uinon-find data structure that can hold at most n
// objects. It uses the quick-find algorithm.
// Returns nil if n < 0.
func NewQuickFind(n int) *QuickFind {
	if n < 0 {
		return nil
	}
	a := make([]int, n, n)
	for i, _ := range a {
		a[i] = i
	}
	return &QuickFind{id: a[:], count: n}
}

// Union implements the UF Union method.
// Between N + 3 to 2N + 1 array access when combining two components.
func (qf *QuickFind) Union(p, q int) {
	if qf.Connected(p, q) {
		return
	}
	// Change all entries with id[p] to id[q]
	pid := qf.id[p]
	qid := qf.id[q]
	for i := 0; i < len(qf.id); i++ {
		if qf.id[i] == pid {
			qf.id[i] = qid
		}
	}
	qf.count--
}

// Find implements the UF Find method.
// Only one array access. (That's why it is called quick-find.)
func (qf *QuickFind) Find(p int) int {
	return qf.id[p]
}

// Connected implements the UF Connected method.
func (qf *QuickFind) Connected(p, q int) bool {
	return qf.id[p] == qf.id[q]
}

// Count implements the UF Count method.
func (qf *QuickFind) Count() int {
	return qf.count
}

// QuickUnion implements the quick-union algorithm.
type QuickUnion struct {
	id    []int // id[i] = parent of node i, node i is root if id[i] == i
	count int   // number of components
}

// NewQuickUnion creates a uinon-find data structure that can hold at most n
// objects. It uses the quick-union algorithm.
// Returns nil if n < 0.
func NewQuickUnion(n int) *QuickUnion {
	if n < 0 {
		return nil
	}
	a := make([]int, n, n)
	for i, _ := range a {
		a[i] = i
	}
	return &QuickUnion{id: a[:], count: n}
}

// Union implements the UF Union method.
// Worst case: 2N - 1 array access.
func (qu *QuickUnion) Union(p, q int) {
	i := qu.Find(p)
	j := qu.Find(q)
	if i == j {
		return
	}
	qu.id[i] = j
	qu.count--
}

// Find implements the UF Find method.
// Cost are related to tree depth. Worst case: N array access.
func (qu *QuickUnion) Find(p int) int {
	// Find the root of node p.
	for p != qu.id[p] {
		p = qu.id[p]
	}
	return p
}

// Connected implements the UF Connected method.
func (qu *QuickUnion) Connected(p, q int) bool {
	return qu.Find(p) == qu.Find(q)
}

// Count implements the UF Count method.
func (qu *QuickUnion) Count() int {
	return qu.count
}

// WeightedQuickUion implements the weighted quick-union algorithm.
type WeightedQuickUnion struct {
	qu   *QuickUnion // Hide implementation detail, so not using embedding here.
	size []int       // size[i] = number of objects in the subtree rooted at i
}

// NewWeightedQuickFind creates a uinon-find data structure that can hold at
// most n objects. It uses the weighted quick-find algorithm. The tree height
// is guaranteed to be at most lg n.
// Returns nil if n < 0.
func NewWeightedQuickUnion(n int) *WeightedQuickUnion {
	qu := NewQuickUnion(n)
	if qu == nil {
		return nil
	}
	size := make([]int, n, n)
	for i := 0; i < n; i++ {
		size[i] = 1
	}
	return &WeightedQuickUnion{qu, size}
}

// Union implements the UF Union method.
// At most ~2(lg n) array accesses.
func (wqu *WeightedQuickUnion) Union(p, q int) {
	// fmt.Println("weighted union called")
	i := wqu.qu.Find(p)
	j := wqu.qu.Find(q)
	if i == j {
		return
	}

	// Make smaller root point to larger one
	if wqu.size[i] < wqu.size[j] {
		wqu.qu.id[i] = j
		wqu.size[j] += wqu.size[i]
	} else {
		wqu.qu.id[j] = i
		wqu.size[i] += wqu.size[j]
	}
	wqu.qu.count--
}

// Find implements the UF Find method. At most lg n array access.
func (wqu *WeightedQuickUnion) Find(p int) int {
	return wqu.qu.Find(p)
}

// Connected implements the UF Connected method.
func (wqu *WeightedQuickUnion) Connected(p, q int) bool {
	return wqu.qu.Connected(p, q)
}

// Count implements the UF Count method.
func (wqu *WeightedQuickUnion) Count() int {
	return wqu.qu.count
}
