package collections

type Iterator[T Object[any]] interface {
	HasNext() bool

	Next() *T
}
