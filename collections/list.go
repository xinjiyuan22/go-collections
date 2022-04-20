package collections

type List[T Object] interface {
	Collection[T]

	Get(index int) *T

	Insert(index int, o T)

	Set(index int, o T)

	SubList(from, to int) List[T]
}
