package set

type Entry interface {
	Compare(another Entry) int
}
