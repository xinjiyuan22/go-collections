package maps

import "github.com/xinjiyuan22/collections/collections"

type HashedMap[K collections.Hashable, V collections.Object] struct {
}

func NewHashMap[K collections.Hashable, V collections.Object]() collections.Map[K, V] {
	return &HashedMap[K, V]{}
}

// Clear: Removes all of the mappings from this map (optional operation).
func (m *HashedMap[K, V]) Clear() {
}

// ContainsKey: Returns true if this map contains a mapping for the specified key.
func (m *HashedMap[K, V]) ContainsKey(k K) bool {
	return true
}

// Size: Returns the number of key-value mappings in this map.
func (m *HashedMap[K, V]) Size() int {
	return 0
}

// Value: Returns a Collection view of the values contained in this map.
func (m *HashedMap[K, V]) Values() collections.Collection[V] {
	return nil
}

// Keys: Returns a Set view of the keys contained in this map.
func (m *HashedMap[K, V]) Keys() collections.Collection[K] {
	return nil
}

// Put: Associates the specified value with the specified key in this map (optional operation).
func (m *HashedMap[K, V]) Put(key K, value V) V {
	return value
}

// Get: Returns the value to which the specified key is mapped, or null if this map contains no mapping for the key.
func (m *HashedMap[K, V]) Get(key K) *V {
	return nil
}

// IsEmpty: Returns true if this map contains no key-value mappings.
func (m *HashedMap[key, value]) IsEmpty() bool {
	return false
}

func (m *HashedMap[K, V]) ToMap() map[K]V {
	return nil
}
