package collections

type Object interface {
	comparable

	Equal(o Object) bool
}
