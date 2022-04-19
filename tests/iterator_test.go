package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xinjiyuan22/collections/list"
	"github.com/xinjiyuan22/collections/objects"
)

func TestArrayIterator(t *testing.T) {
	array := list.NewArrayListFromSlice([]*objects.Integer{objects.IntegerPtr(1), objects.IntegerPtr(2), objects.IntegerPtr(3)})
	iter := array.Iterator()
	assert.True(t, iter.HasNext())
	assert.True(t, iter.Next().Equal(objects.Integer(1)))

	assert.True(t, iter.HasNext())
	assert.True(t, iter.Next().Equal(objects.Integer(2)))

	assert.True(t, iter.HasNext())
	assert.True(t, iter.Next().Equal(objects.Integer(3)))

	assert.False(t, iter.HasNext())
	assert.Nil(t, iter.Next())

	array2 := list.NewArrayListFromSlice([]*objects.Integer{})
	iter = array2.Iterator()
	assert.False(t, iter.HasNext())
	assert.Nil(t, iter.Next())
}
