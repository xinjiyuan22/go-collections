package collections

type Hashable interface {
	Object

	comparable

	HashCode() int64
}
