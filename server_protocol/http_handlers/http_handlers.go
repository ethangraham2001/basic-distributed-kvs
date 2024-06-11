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
const APIEndpoint string = "/"

const contentType string = "Content-Type"
const contentTypeOctetStream string = "application/octet-stream"
const contentTypeJson string = "application/json"

const maxBytes int = 10 << 20; // 10MiB limit
 
// defines the put request body for json unmarshalling
type putRequest[K comparable] struct {
	Key K `json:"key"`
}

// defines the get request body for json unmarshalling
type getRequest[K comparable] struct {
	Key K `json:"key"`
}

// InitHandleReq initializes a HandlerFunc that handles get and put requests
// for a Peer p.
// Get requests will return a raw byte array, which is why we contrain the
// Peer.
func InitHandleReq[K comparable](p *peer.Peer[K, []byte]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handleGet(w, r, p)
		case http.MethodPut:
			handlePut(w, r, p)
		}
	}
}

func handleGet[K comparable](w http.ResponseWriter, r *http.Request, p *peer.Peer[K, []byte]) {
    log.Print("Received GET request")
	var decodedReq getRequest[K]

	err := json.NewDecoder(r.Body).Decode(&decodedReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	key := decodedReq.Key

	data, err := p.GetFromDatastore(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	w.Header().Set(contentType, contentTypeOctetStream)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func handlePut[K comparable](w http.ResponseWriter, r *http.Request, p *peer.Peer[K, []byte]) {
    log.Print("Recived PUT request")
	var decodedReq putRequest[K]

	err := json.NewDecoder(r.Body).Decode(&decodedReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

    reader, err := r.MultipartReader()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    var key K
    var data []byte

    for {
        part, err := reader.NextPart()

        if err == io.EOF {
            break;
        }

        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        if part.Header.Get(contentType) == contentTypeJson {
            key, err = decodePutJson[K](r)
            if err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
                return
            }
        }

        if part.Header.Get(contentType) == contentTypeOctetStream {
            var data []byte
            bytesRead, err := part.Read(data)
            
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }

            if bytesRead == 0 {
                http.Error(w, "Empty file", http.StatusBadRequest)
                return
            }
        }
    }

    p.PutInDataStore(key, data)
}

// decodes put request json and returns the key
func decodePutJson[K comparable](r *http.Request) (K, error) {
    var decodedReq putRequest[K]
    var key K

    err := json.NewDecoder(r.Body).Decode(&decodedReq)
	if err != nil {
        return key, err
	}

    return key, nil
}

