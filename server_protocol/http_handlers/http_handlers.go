package httphandlers

import (
	"net/http"

	"github.com/ethangraham2001/distributed_kvs_server_protocol/server"
)

var srv server.Server[int64, string]

type PutRequest struct {
	key   int64
	value string
}

type GetRequest struct {
	key int64
}

func HandleReq(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGet(w, r)
	case http.MethodPut:
		handlePut(w, r)
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) {
}

func handlePut(w http.ResponseWriter, r *http.Request) {
}
