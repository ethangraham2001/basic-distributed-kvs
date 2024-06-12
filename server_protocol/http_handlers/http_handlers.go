// Package httphandlers defines http handlers for client-to-peer communication
// or peer-to-peer communication. Any business logic related to keys that
// are contained in different Peers is not handled here
package httphandlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/ethangraham2001/distributed_kvs_server_protocol/peer"
)

// APIEndpoint is the only available Endpoint for now.
const APIEndpoint string = "/api/"

const contentType string = "Content-Type"
const contentTypeOctetStream string = "application/octet-stream"
const contentTypeJSON string = "application/json"

const maxBytes int = 10 << 20 // 10MiB limit

// defines the put request body for json unmarshalling
type putRequest[K comparable] struct {
	Key K `json:"key"`
}

// defines the get request body for json unmarshalling
type getRequest[K comparable] struct {
	Key K `json:"key"`
}

// InitHandleReq initializes and returns a HandlerFunc that handles get
// and put requests for a Peer p.
// Get requests will return a raw byte array, which is why we contrain the
// Peer.
func InitHandleReq(p *peer.Peer[string, []byte]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handleGet(w, r, p)
		case http.MethodPut:
			handlePut(w, r, p)
		}
	}
}

func handleGet(w http.ResponseWriter, r *http.Request, p *peer.Peer[string, []byte]) {
	key := getKeyFromURLPath(r.URL.Path)
	log.Print("Received GET request")

	data, err := p.GetFromDatastore(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	w.Header().Set(contentType, contentTypeOctetStream)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func handlePut(w http.ResponseWriter, r *http.Request, p *peer.Peer[string, []byte]) {
	if r.Header.Get(contentType) != contentTypeOctetStream {
		http.Error(w, "Require octet stream", http.StatusUnsupportedMediaType)
		return
	}

	log.Print("received PUT request")
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer r.Body.Close()

	key := getKeyFromURLPath(r.URL.Path)
	p.PutInDataStore(key, data)
	w.WriteHeader(http.StatusOK)
}

// decodes put request json and returns the key
func decodePutJSON[K comparable](r *http.Request) (K, error) {
	var decodedReq putRequest[K]
	var key K

	err := json.NewDecoder(r.Body).Decode(&decodedReq)
	if err != nil {
		return key, err
	}

	return key, nil
}

func getKeyFromURLPath(url string) string {
	return url[len(APIEndpoint):]
}
