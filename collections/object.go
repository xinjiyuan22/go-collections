package collections

type Object[T any] interface {
	Equal(o Object[T]) bool

	Less(o Object[T]) bool
}
