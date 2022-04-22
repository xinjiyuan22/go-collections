package collections

type Comparable interface {
	Object

	Less(x Comparable) bool
}
