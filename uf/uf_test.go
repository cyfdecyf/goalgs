package uf

import (
	"testing"
)

type pair struct {
	p int
	q int
}

var tinyUF = [...]pair{
	{4, 3},
	{3, 8},
	{6, 5},
	{9, 4},
	{2, 1},
	{8, 9},
	{5, 0},
	{7, 2},
	{6, 1},
	{1, 0},
	{6, 7},
}

func testUF(uf UF, t *testing.T, ufname string) {
	for _, s := range tinyUF {
		uf.Union(s.p, s.q)
	}

	if uf.Connected(1, 2) != true {
		t.Error(ufname, "1 and 2 should be connected ")
	}

	if uf.Connected(1, 7) != true {
		t.Error(ufname, "1 and 7 should be connected")
	}

	if uf.Connected(9, 6) == true {
		t.Error(ufname, "1 and 9 should NOT be connected")
	}

	cnt := uf.Count()
	if cnt != 3 {
		t.Error(ufname, "component should be 3, got:", cnt)
	}
}

func TestQuickFind(t *testing.T) {
	uf := NewQuickFind(len(tinyUF))
	testUF(uf, t, "QuickFind")
}

func TestQuickUnion(t *testing.T) {
	uf := NewQuickUnion(len(tinyUF))
	testUF(uf, t, "QuickUnion")
}

func TestWeightedQuickUnion(t *testing.T) {
	uf := NewWeightedQuickUnion(len(tinyUF))
	testUF(uf, t, "WeightedQuickUnion")
}
