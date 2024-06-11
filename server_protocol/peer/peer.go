// Package peer provides type definition of Peer struct, as well as relevant
// methods used for communication
package peer

import (
	"github.com/ethangraham2001/distributed_kvs_server_protocol/datastore"
)

// Peer in the network. Must be configured at program launch.
type Peer[K comparable, V any] struct {
	// ID is a unique identifier for the peer in the network.
	// This peer will store data such that hash(key) mod n = ID,
	// with n the number of peers in the network.
	ID uint32
	// datastore is the primary source of data that the peer
	// fetches data from.
	datastore datastore.Datastore[K, V]
	// connections maps identifiers to IP addresses for communication
	// with other peers in the network.
	// The Peer contains its own mapping within connections.
	connections map[uint32]Address
}

// GetFromDatastore returns the value with key `key` from the Peer's
// datastore
func (p *Peer[K, V]) GetFromDatastore(key K) (V, error) {
	value, err := p.datastore.Get(key)
	return value, err
}

// PutInDataStore puts `(key, value)` into the Peer's datastore
func (p *Peer[K, V]) PutInDataStore(key K, value V) {
	p.datastore.Put(key, value)
}

// NewPeer initializes a new Peer from a configuration file
func newPeer[K comparable, V any](ID uint32) Peer[K, V] {
	d := datastore.NewInMemoryDataStore[K, V]()
	return Peer[K, V]{
		ID:          ID,
		datastore:   d,
		connections: make(map[uint32]Address),
	}
}

func (p *Peer[K, V]) addConnection(id uint32, addr Address) {
	p.connections[id] = addr
}
