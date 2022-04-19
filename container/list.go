package container

type List[T Object[any]] interface {
	Container

	Add(index int, o T)

	Append(o ...T)
}
