package set

import "github.com/itsabgr/go-set/pkg/list"

type iter[T Entry[T]] struct {
	e *list.Element[T]
}

func (i *iter[T]) Item() T {
	return i.e.Value
}

func (i *iter[T]) HasNext() bool {
	return i.e != nil
}

func (i *iter[T]) Next() {
	i.e = i.e.Next()
}
