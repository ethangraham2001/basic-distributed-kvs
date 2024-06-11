// Package main is the main package of the project.
package main

import (
	"log"
	"net/http"

	"github.com/ethangraham2001/distributed_kvs_server_protocol/http_handlers"
	"github.com/ethangraham2001/distributed_kvs_server_protocol/peer"
)

const filepath string = "config.yaml"

func main() {
	p, err := peer.ReadConfigFromFile[string, []byte](filepath)
	if err != nil {
		panic("failed to initialize peer!")
	}

	handlerFunc := httphandlers.InitHandleReq(&p)
	http.HandleFunc(httphandlers.APIEndpoint, handlerFunc)
	log.Print("Launching peer")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
