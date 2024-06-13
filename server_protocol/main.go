// Package main is the main package of the project.
package main

import (
	"fmt"
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
	portNum := p.GetConnection(p.ID).Port
	portStr := fmt.Sprintf(":%d", portNum)
	http.HandleFunc(peer.APIEndpoint, handlerFunc)
	log.Printf("Peer listening @ localhost:%d", portNum)
	log.Fatal(http.ListenAndServe(portStr, nil))
}
