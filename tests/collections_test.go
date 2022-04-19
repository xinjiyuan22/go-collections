package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/xinjiyuan22/collections/collections"
	"github.com/xinjiyuan22/collections/list"
	"github.com/xinjiyuan22/collections/objects"
)

func testCollectionAdd(t *testing.T, c collections.Collection[objects.Integer]) {
	c.Clear()
	assert.Equal(t, c.Add(objects.Integer(1)), true)
	assert.Equal(t, c.Add(objects.Integer(2)), true)
	assert.Equal(t, c.ToArray(), []objects.Integer{1, 2})
}

func testCollectionAddAll(t *testing.T, c collections.Collection[objects.Integer]) {
	c.Clear()
	assert.True(t, c.Add(objects.Integer(1)))
	add := list.NewArrayList[objects.Integer](1)
	add.Add(1)
	add.Add(2)

	assert.True(t, c.AddAll(add))
	assert.Equal(t, c.ToArray(), []objects.Integer{1, 1, 2})
}

func testCollectionClear(t *testing.T, c collections.Collection[objects.Integer]) {
	c.Clear()
	c.Add(1)
	c.Add(2)
	assert.Equal(t, []objects.Integer{1, 2}, c.ToArray())
	c.Clear()
	assert.True(t, c.IsEmpty())
}

func testCollectionContains(t *testing.T, c collections.Collection[objects.Integer]) {
	c.Clear()

	c.Add(1)
	c.Add(2)
	c.Add(1)

	assert.True(t, c.Contains(1))
	assert.True(t, c.Contains(2))
	assert.False(t, c.Contains(3))
}

func testCollectionContainsAll(t *testing.T, c collections.Collection[objects.Integer]) {
	c.Clear()

	c.Add(1)
	c.Add(2)
	c.Add(3)

	array1 := list.NewArrayListFromSlice([]*objects.Integer{objects.IntegerPtr(1), objects.IntegerPtr(2), objects.IntegerPtr(3)})
	assert.True(t, c.ContainsAll(array1))

	array2 := list.NewArrayListFromSlice([]*objects.Integer{objects.IntegerPtr(1), objects.IntegerPtr(3)})
	assert.True(t, c.ContainsAll(array2))

	array3 := list.NewArrayListFromSlice([]*objects.Integer{objects.IntegerPtr(1), objects.IntegerPtr(4)})
	assert.False(t, c.ContainsAll(array3))

	array4 := list.NewArrayListFromSlice([]*objects.Integer{objects.IntegerPtr(5), objects.IntegerPtr(4)})
	assert.False(t, c.ContainsAll(array4))
}

func testCollectionIsEmpty(t *testing.T, c collections.Collection[objects.Integer]) {
	c.Clear()
	assert.True(t, c.IsEmpty())
	c.Add(1)
	assert.False(t, c.IsEmpty())
}

func testCollectionRemove(t *testing.T, c collections.Collection[objects.Integer]) {
	c.Clear()
	c.Add(1)
	c.Add(2)
	c.Add(3)
	c.Add(1)
	c.Add(5)

	assert.True(t, c.Remove(1))
	assert.Equal(t, []objects.Integer{2, 3, 5}, c.ToArray())
	assert.True(t, c.Remove(2))
	assert.Equal(t, []objects.Integer{3, 5}, c.ToArray())

	c.Clear()
	c.Add(1)
	c.Add(2)
	c.Add(3)
	c.Add(1)
	c.Add(5)
	assert.False(t, c.Remove(6))
	assert.Equal(t, []objects.Integer{1, 2, 3, 1, 5}, c.ToArray())

	c.Clear()
	c.Add(1)
	c.Add(2)
	c.Add(3)
	c.Add(1)
	c.Add(5)
	assert.True(t, c.Remove(5))
	assert.Equal(t, []objects.Integer{1, 2, 3, 1}, c.ToArray())

	c.Clear()
	assert.False(t, c.Remove(1))
	assert.True(t, c.IsEmpty())
}

func testCollectionRemoveAll(t *testing.T, c collections.Collection[objects.Integer]) {
	c.Clear()
	c.Add(1)
	c.Add(2)
	c.Add(3)
	c.Add(1)
	c.Add(5)
	array1 := list.NewArrayListFromSlice([]*objects.Integer{objects.IntegerPtr(1), objects.IntegerPtr(2), objects.IntegerPtr(4)})
	assert.True(t, c.RemoveAll(array1))
	assert.Equal(t, []objects.Integer{3, 5}, c.ToArray())

	c.Clear()
	c.Add(1)
	c.Add(2)
	c.Add(3)
	c.Add(1)
	c.Add(5)
	array2 := list.NewArrayListFromSlice([]*objects.Integer{objects.IntegerPtr(3), objects.IntegerPtr(3), objects.IntegerPtr(5)})
	assert.True(t, c.RemoveAll(array2))
	assert.Equal(t, []objects.Integer{1, 2, 1}, c.ToArray())

	c.Clear()
	c.Add(1)
	c.Add(2)
	c.Add(3)
	c.Add(1)
	c.Add(5)
	array3 := list.NewArrayListFromSlice([]*objects.Integer{objects.IntegerPtr(4), objects.IntegerPtr(6), objects.IntegerPtr(7)})
	assert.False(t, c.RemoveAll(array3))
	assert.Equal(t, []objects.Integer{1, 2, 3, 1, 5}, c.ToArray())

	c.Clear()
	c.Add(1)
	c.Add(2)
	c.Add(3)
	c.Add(1)
	c.Add(5)
	array4 := list.NewArrayListFromSlice([]*objects.Integer{})
	assert.False(t, c.RemoveAll(array4))
	assert.Equal(t, []objects.Integer{1, 2, 3, 1, 5}, c.ToArray())
}

func testCollectionSize(t *testing.T, c collections.Collection[objects.Integer]) {
	c.Clear()
	assert.Equal(t, 0, c.Size())

	c.Add(1)
	assert.Equal(t, 1, c.Size())
}

func testCollection(t *testing.T, c collections.Collection[objects.Integer]) {
	t.Log("\n=========== Testing Add Method ===========\n")
	testCollectionAdd(t, c)

	t.Log("\n=========== Testing AddAll Method ===========\n")
	testCollectionAddAll(t, c)

	t.Log("\n=========== Testing Clear Method ===========\n")
	testCollectionClear(t, c)

	t.Log("\n=========== Testing Contains Method ===========\n")
	testCollectionContains(t, c)

	t.Log("\n=========== Testing ContainsAll Method ===========\n")
	testCollectionContainsAll(t, c)

	t.Log("\n=========== Testing IsEmpty Method ===========\n")
	testCollectionIsEmpty(t, c)

	t.Log("\n=========== Testing Remove Method ===========\n")
	testCollectionRemove(t, c)

	t.Log("\n=========== Testing RemoveAll Method ===========\n")
	testCollectionRemoveAll(t, c)

	t.Log("\n=========== Testing Size Method ===========\n")
	testCollectionSize(t, c)
}

func TestArrayList(t *testing.T) {
	array := list.NewArrayList[objects.Integer](10)
	testCollection(t, array)
}
