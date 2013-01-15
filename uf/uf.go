// Package uf implements 3 union-find algorithms: quick-find, quick-union and TODO

package uf

import (
// "fmt"
)

// Union-find API
type UF interface {
	// Add connection between p and q.
	Union(p, q int)
	// Return component identifier for p.
	Find(p int) int
	// Return true if p and q are in the same component
	Connected(p, q int) bool
	// Number of components
	Count() int
}

// Quick-find algorithm.
type QuickFind struct {
	id  []int // p and q are connected iff they have the same id in the array
	cnt int   // number of components
}

// Creat a QuickFind instance which implements the UF interface.
// Returns nil if n < 0.
func NewQuickFind(n int) *QuickFind {
	if n < 0 {
		return nil
	}
	a := make([]int, n, n)
	for i, _ := range a {
		a[i] = i
	}
	return &QuickFind{id: a[:], cnt: n}
}

// Find implements the UF Find method.
// Only one array access.
func (qf *QuickFind) Find(p int) int {
	return qf.id[p]
}

// Connected implements the UF Connected method.
func (qf *QuickFind) Connected(p, q int) bool {
	return qf.id[p] == qf.id[q]
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
	qf.cnt--
}

// Count implements the UF Count method.
func (qf *QuickFind) Count() int {
	return qf.cnt
}

// Quick-union algorithm
type QuickUnion struct {
	id  []int // id[p] contains the parent of node p, node p is root if id[p] == p
	cnt int   // number of components
}

// Count implements the UF Count method.
func NewQuickUnion(n int) *QuickUnion {
	if n < 0 {
		return nil
	}
	a := make([]int, n, n)
	for i, _ := range a {
		a[i] = i
	}
	return &QuickUnion{id: a[:], cnt: n}
}

// Find implements the UF Find method.
// Worst case: N array access.
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

// Union implements the UF Union method.
// Worst case: 2N - 1 array access.
func (qu *QuickUnion) Union(p, q int) {
	i := qu.Find(p)
	j := qu.Find(q)
	if i == j {
		return
	}
	qu.id[i] = j
	qu.cnt--
}

// Count implements the UF Count method.
func (qu *QuickUnion) Count() int {
	return qu.cnt
}
