package collections

type Collection[T Object] interface {

	// Add: Ensures that this collection contains the specified element (optional operation).
	Add(o T) bool

	AddAll(c Collection[T]) bool

	Clear()

	Contains(o T) bool

	ContainsAll(c Collection[T]) bool

	IsEmpty() bool

	Iterator() Iterator[T]

	Remove(o T) bool

	RemoveAll(c Collection[T]) bool

	ToArray() []T

	Size() int
}
