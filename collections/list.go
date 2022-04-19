package collections

type List[T Object[any]] interface {
	Collection[T]

	Get(index int) *T

	Append(o ...T)

	Set(index int, o T)

	SubList(from, to int) List[T]
}
