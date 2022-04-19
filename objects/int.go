package objects

import "github.com/xinjiyuan22/collections/collections"

type Integer int

func (i Integer) Equal(v collections.Object) bool {
	if value, ok := v.(Integer); !ok {
		return false
	} else {
		return value == i
	}
}

func IntegerPtr(v Integer) *Integer {
	return &v
}
