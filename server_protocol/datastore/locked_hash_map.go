package datastore

import "sync"

// defines a thread-safe hashmap for types key type `K` and value type `V`
type LockedHashMap[K comparable, V any] struct {
	mutex sync.Mutex
	data  map[K]V
}

func (self *LockedHashMap[K, V]) Get(key K) (V, bool) {
	self.mutex.Lock()
	defer self.mutex.Unlock()

	value, contains := self.data[key]
	if !contains {
		return value, contains
	}

	return value, true
}

func (self *LockedHashMap[K, V]) Put(key K, value V) {
	self.mutex.Lock()
	defer self.mutex.Unlock()

	self.data[key] = value
}

func NewLockedHashMap[K comparable, V any]() LockedHashMap[K, V] {
	data := make(map[K]V)
	return LockedHashMap[K, V]{
		data: data,
	}
}
