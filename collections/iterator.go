package collections

type Iterator[T Object] interface {
	HasNext() bool

	Next() *T
}
