package collections

type Map[KEY Hashable, VALUE Object] interface {
	// Clear: Removes all of the mappings from this map (optional operation).
	Clear()

	// ContainsKey: Returns true if this map contains a mapping for the specified key.
	ContainsKey(key KEY) bool

	// Size: Returns the number of key-value mappings in this map.
	Size() int

	// Value: Returns a Collection view of the values contained in this map.
	Values() Collection[VALUE]

	// Keys: Returns a Set view of the keys contained in this map.
	Keys() Collection[KEY]

	// Put: Associates the specified value with the specified key in this map (optional operation).
	Put(key KEY, value VALUE) VALUE

	// Get: Returns the value to which the specified key is mapped, or null if this map contains no mapping for the key.
	Get(key KEY) *VALUE

	// IsEmpty: Returns true if this map contains no key-value mappings.
	IsEmpty() bool

	ToMap() map[KEY]VALUE
}
