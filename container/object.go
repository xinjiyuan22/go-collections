package container

type Object[T any] interface {
	Equal(o T) bool

	Less(o T) bool
}
