package set

import (
	"github.com/itsabgr/go-set/pkg/list"
)

type Set[T Entry[T]] struct {
	items list.List[T]
}

func (set *Set[T]) Init() *Set[T] {
	set.items.Init()
	return set
}

func NewSet[T Entry[T]]() *Set[T] {
	return (&Set[T]{}).Init()
}

func (set *Set[T]) Has(item T) bool {
	for e := set.items.Front(); e != nil; e = e.Next() {
		if e.Value.Compare(item) == 0 {
			return true
		}
	}
	return false
}

func (set *Set[T]) Add(item T) bool {
	if set.Has(item) {
		return false
	}
	set.items.PushBack(item)
	return true
}

func (set *Set[T]) Len() int {
	return set.items.Len()
}

func (set *Set[T]) Iter() *iter[T] {
	i := &iter[T]{e: set.items.Front()}
	return i
}

func (set *Set[T]) Remove(item T) bool {
	for e := set.items.Front(); e != nil; e = e.Next() {
		if e.Value.Compare(item) == 0 {
			set.items.Remove(e)
			return true
		}
	}
	return false
}
