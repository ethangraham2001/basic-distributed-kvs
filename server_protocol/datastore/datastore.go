package datastore

// interface for a datastore
type Datastore[K any, V any] interface {
	// returns the value at `key`, or an error. Thread-safe
	Get(key K) (V, error)

	// puts `value` at `key`, or returns an error if unsuccessful. Thread-safe
	Put(key K, value V)
}
