// Package main is the main package of the client protocol
package main

import (
	"log"
	"os"

	"github.com/ethangraham2001/distributed_kvs_client_protocol/client"
	"github.com/ethangraham2001/distributed_kvs_client_protocol/address"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("please provide key as well as filepath")
	}

    client := client.NewClient[string]()
	peerID := uint32(0)
    peerAddr := address.NewAddress("127.0.0.1", 8080)

    client.AddConnection(peerID, peerAddr)

    data := []byte("Hello, world!")
    
    log.Print(data)

    if os.Args[1] == "PUT" {
        err := client.MakePutRequest(peerID, "key", data)
        if err != nil {
            log.Fatal(err)
        }
    } else if os.Args[1] == "GET" {
        data, err := client.MakeGetRequest(peerID, "key")
        if err != nil {
            log.Fatal(err)
        }
        log.Printf("Received data = %s\n", string(data))
    } else {
        log.Fatal("Invalid argument")
    }
}
