package collections

type Hashable interface {
	Object

	HashCode() int64
}
