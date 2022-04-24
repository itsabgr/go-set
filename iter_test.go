package set

import (
	"github.com/itsabgr/go-handy"
	"sort"
	"testing"
)

type ent int

func (e ent) Compare(e2 ent) int {
	if e == e2 {
		return 0
	}
	return -1
}

func TestIter(t *testing.T) {
	set := NewSet[ent]()
	var items []int
	for i := range handy.N(100) {
		items = append(items, i)
		if !set.Add(ent(i)) {
			t.Fatal()
		}
	}
	for i := range handy.N(100) {
		if set.Add(ent(i)) {
			t.Fatal()
		}
	}
	var items2 []int
	for iter := set.Iter(); iter.HasNext(); iter.Next() {
		items2 = append(items2, int(iter.Item()))
	}
	sort.Ints(items2)
	sort.Ints(items)
	for i := range items2 {
		if items[i] != items2[i] {
			t.Fatal()
		}
	}
}
