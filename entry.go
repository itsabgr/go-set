package set

type Entry[E any] interface {
	Compare(another E) int
}
