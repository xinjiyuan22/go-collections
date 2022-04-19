package collections

type HashObject[T any] interface {
	Object[T]

	Hash() string
}
