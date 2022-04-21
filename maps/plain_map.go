package maps

import (
	"sync"

	"github.com/xinjiyuan22/collections/collections"
	"github.com/xinjiyuan22/collections/list"
)

type PlainMap[K collections.Hashable, V collections.Object] struct {
	sync.RWMutex
	plain  map[K]V
	values collections.Collection[V]
	keys   collections.Collection[K]
}

func NewPlainMap[K collections.Hashable, V collections.Object]() collections.Map[K, V] {
	return &PlainMap[K, V]{
		plain:  make(map[K]V),
		values: list.NewLinkedList[V](),
		keys:   list.NewLinkedList[K](),
	}
}

// Clear: Removes all of the mappings from this map (optional operation).
func (m *PlainMap[K, V]) Clear() {
	m.Lock()
	defer m.Unlock()

	m.plain = make(map[K]V)
}

// ContainsKey: Returns true if this map contains a mapping for the specified key.
func (m *PlainMap[K, V]) ContainsKey(k K) bool {
	_, ok := m.plain[k]
	return ok
}

// Size: Returns the number of key-value mappings in this map.
func (m *PlainMap[K, V]) Size() int {
	return len(m.plain)
}

// Value: Returns a Collection view of the values contained in this map.
func (m *PlainMap[K, V]) Values() collections.Collection[V] {
	return m.values
}

// Keys: Returns a Set view of the keys contained in this map.
func (m *PlainMap[K, V]) Keys() collections.Collection[K] {
	return m.keys
}

// Put: Associates the specified value with the specified key in this map (optional operation).
func (m *PlainMap[K, V]) Put(key K, value V) V {
	m.Lock()
	defer m.Unlock()
	m.plain[key] = value
	m.keys.Add(key)
	m.values.Add(value)
	return value
}

// Get: Returns the value to which the specified key is mapped, or null if this map contains no mapping for the key.
func (m *PlainMap[K, V]) Get(key K) *V {
	m.RLock()
	defer m.RUnlock()
	v, ok := m.plain[key]
	if !ok {
		return nil
	}
	return &v
}

// IsEmpty: Returns true if this map contains no key-value mappings.
func (m *PlainMap[key, value]) IsEmpty() bool {
	return len(m.plain) == 0
}

func (m *PlainMap[K, V]) ToMap() map[K]V {
	return m.plain
}
