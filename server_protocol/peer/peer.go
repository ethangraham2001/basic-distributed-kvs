// Package peer provides type definition of Peer struct, as well as relevant
// methods used for communication
package peer

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"

	"github.com/ethangraham2001/distributed_kvs_server_protocol/datastore"
	"github.com/ethangraham2001/distributed_kvs_server_protocol/util"
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
	// N is the number of peers that a piece of data is replicated to. This
	// is equal to the length of the data's perferences list
	N uint32
}

// newPeer initializes a new Peer from a configuration file
func newPeer[K comparable, V any](ID uint32, N uint32) Peer[K, V] {
	d := datastore.NewInMemoryDataStore[K, V]()
	return Peer[K, V]{
		ID:          ID,
		datastore:   d,
		connections: make(map[uint32]Address),
		N:           N,
	}
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

func (p *Peer[K, V]) addConnection(id uint32, addr Address) {
	p.connections[id] = addr
}

// GetConnection returns the addres of the Peer identified by ID. This can
// be the Peer itself
func (p *Peer[K, V]) GetConnection(id uint32) Address {
	return p.connections[id]
}

// IsLastInPrefList returns `true` iff the peer is the last in the preferences
// list for a given key. Key type is defined as string so that it is hashable
func (p *Peer[K, V]) IsLastInPrefList(key string) bool {
	firstPeer := util.HashKey(key) % uint32(len(p.connections))
	return p.ID == firstPeer+p.N
}

func (p *Peer[K, V]) nextPeerID() uint32 {
	return (p.ID + 1) % uint32(len(p.connections))
}

func (p *Peer[K, V]) ReplicateToNextPeer(key string, data []byte) error {
	addr, valid := p.connections[p.nextPeerID()]
	if !valid {
		return errors.New("Invalid next peer ID")
	}
	path := APIEndpoint + key

	req, err := http.NewRequest(http.MethodPut, addr.String()+path, bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/octet-stream")
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		errorMsg := fmt.Sprintf("PUT failed. Peer returned %s", resp.Status)
		return errors.New(errorMsg)
	}

	return nil
}
