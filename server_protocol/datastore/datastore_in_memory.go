package datastore

import "errors"

// defines a datastore that is contained in-memory
type InMemoryDataStore[K comparable, V any] struct {
	data *LockedHashMap[K, V]
}

func (self InMemoryDataStore[K, V]) Get(key K) (V, error) {
	value, contains := self.data.Get(key)
	if !contains {
		return value, errors.New("No such element")
	}

	return value, nil
}

func (self InMemoryDataStore[K, V]) Put(key K, value V) {
	self.data.Put(key, value)
}

// NewInMemoryDataStore initializes a new in memory datastore and returns it.
func NewInMemoryDataStore[K comparable, V any]() InMemoryDataStore[K, V] {
	data := NewLockedHashMap[K, V]()
	return InMemoryDataStore[K, V]{
		data: &data,
	}
}
