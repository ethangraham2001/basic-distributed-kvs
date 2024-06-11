package server

// this is a test
import (
	"github.com/ethangraham2001/distributed_kvs_server_protocol/datastore"
	"github.com/ethangraham2001/distributed_kvs_server_protocol/util"
)

type keyValuePair[K comparable, V any] struct {
	key   K
	value V
}

// Server: Handles incoming requests in a multi-threaded fashion
type Server[K comparable, V any] struct {
	datastore  datastore.Datastore[K, V]
	keyIn      <-chan K
	valueOut   chan<- util.Result[V]
	keyValueIn chan keyValuePair[K, V]
}

// thread-safe get from datastore. Should be called from goroutine
func (self *Server[K, V]) Get() {
	key := <-self.keyIn
	value, err := self.datastore.Get(key)
	self.valueOut <- util.Result[V]{Value: value, Err: err}
}

// thread-safe put into datastore. Should be called from goroutine
func (self *Server[K, V]) Put() {
	keyValue := <-self.keyValueIn
	key := keyValue.key
	value := keyValue.value

	self.datastore.Put(key, value)
}
