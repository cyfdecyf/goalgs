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

// Eager approach implementing the UF API.
type QuickFind struct {
	id  []int // p and q are connected iff they have the same id in the array
	cnt int   // number of components
}

// Creat a QuickFind instance which implements the UF interface.
func NewQuickFind(n int) *QuickFind {
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
	// Change all entries with id[p] to id[q]
	pid := qf.id[p]
	qid := qf.id[q]

	// Nothing to do if p and q are already in the same component.
	if pid == qid {
		return
	}

	// Rename p’s component to q’s name.
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
