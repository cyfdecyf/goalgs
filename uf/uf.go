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
	qf := QuickFind{id: make([]int, n, n), count: n}
	for i := 0; i < n; i++ {
		qf.id[i] = i
	}
	return &qf
}

// Union implements the UF Union method.
// Between N + 3 to 2N + 1 array access when combining two components.
func (qf *QuickFind) Union(p, q int) {
	if qf.Connected(p, q) {
		return
	}
	// change all entries with id[p] to id[q]
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
	qu := QuickUnion{id: make([]int, n, n), count: n}
	for i := 0; i < n; i++ {
		qu.id[i] = i
	}
	return &qu
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
	// find the root of node p
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

// WeightedQuickUion implements the weighted quick-union algorithm (with one
// pass path compression).
type WeightedQuickUnion struct {
	id    []int // id[i] = parent of node i, node i is root if id[i] == i
	size  []int // size[i] = number of objects in the subtree rooted at i
	count int
}

// NewWeightedQuickFind creates a uinon-find data structure that can hold at
// most n objects. It uses the weighted quick-find algorithm. The tree height
// is guaranteed to be at most lg n. (With path compression, the tree is
// almost completely flat.)
// Returns nil if n < 0.
func NewWeightedQuickUnion(n int) *WeightedQuickUnion {
	if n < 0 {
		return nil
	}
	wqu := WeightedQuickUnion{id: make([]int, n, n),
		size: make([]int, n, n), count: n}
	for i := 0; i < n; i++ {
		wqu.id[i] = i
		wqu.size[i] = 1
	}
	return &wqu
}

// Union implements the UF Union method.
// At most ~2(lg n) array accesses.
func (wqu *WeightedQuickUnion) Union(p, q int) {
	// fmt.Println("weighted union called")
	i := wqu.Find(p)
	j := wqu.Find(q)
	if i == j {
		return
	}

	// make smaller root point to larger one
	if wqu.size[i] < wqu.size[j] {
		wqu.id[i] = j
		wqu.size[j] += wqu.size[i]
	} else {
		wqu.id[j] = i
		wqu.size[i] += wqu.size[j]
	}
	wqu.count--
}

// Find implements the UF Find method. At most lg n array access.
func (wqu *WeightedQuickUnion) Find(p int) int {
	for p != wqu.id[p] {
		// path compression: point to grandparent if not root
		wqu.id[p] = wqu.id[wqu.id[p]]
		p = wqu.id[p]
	}
	return p
}

// Connected implements the UF Connected method.
func (wqu *WeightedQuickUnion) Connected(p, q int) bool {
	return wqu.Find(p) == wqu.Find(q)
}

// Count implements the UF Count method.
func (wqu *WeightedQuickUnion) Count() int {
	return wqu.count
}
