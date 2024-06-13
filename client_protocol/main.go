// Package main is the main package of the client protocol
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethangraham2001/distributed_kvs_client_protocol/client"
	"github.com/ethangraham2001/distributed_kvs_client_protocol/util"
)

const usageStr string = `
Usage:
    ./distributed_kvs_client_protocol GET <key> <output_filename>
    ./distributed_kvs_client_protocol PUT <key> <input_filename>
`

func usage() {
	log.Fatal(usageStr)
}

func main() {
	if len(os.Args) != 4 {
		usage()
	}
	fmt.Println(os.Args)

	client, err := client.ReadConfigFromFile("config.yaml")
	if err != nil {
		log.Fatalf("Failed to read from configuration file: %s", err.Error())
	}

	if os.Args[1] == "PUT" {
		handlePut(&client, os.Args[2], os.Args[3])
	} else if os.Args[1] == "GET" {
		handleGet(&client, os.Args[2], os.Args[3])
	} else {
		usage()
	}
}

// handlePut attempts to a piece of data in the first peer from the
// preferences list, making its way down upon failure until success or
// exhaustion of the list
func handlePut(c *client.Client, key string, filePath string) {
	data, err := util.ReadFile(filePath)
	if err != nil {
		log.Fatal(err.Error())
	}

	keyHash := util.HashKey(key)
	var peerID uint32

	for i := uint32(0); i <= c.N; i++ {
		peerID = (keyHash + i) % c.NumPeers
		err = c.MakePutRequest(peerID, key, data)
		if err != nil {
			log.Printf("Failed to write to Peer_%d", peerID)
		} else {
			goto success
		}
	}
	log.Fatalf("Failed to make PUT request. Peers unresponsive. %s", err.Error())

success:
	log.Printf("Data stored in Peer_%d", peerID)
}

// handleGet attempts to put get a piece of data in the first peer from the
// preferences list, making its way down upon failure until success or
// exhaustion of the list
func handleGet(c *client.Client, key string, filePath string) {
	keyHash := util.HashKey(key)

	var (
		data   []byte
		err    error
		peerID uint32
	)
	for i := uint32(0); i <= c.N; i++ {
		peerID = (keyHash + i) % c.NumPeers
		data, err = c.MakeGetRequest(peerID, key)
		if err != nil {
			log.Printf("Failed to query Peer_%d", peerID)
		} else {
			goto write_data
		}
	}
	log.Fatalf("Failed to make GET request. Peers unresponsive. %s", err.Error())
	return

write_data:
	log.Printf("Data retrieved from Peer_%d", peerID)
	err = util.WriteFile(filePath, data)
	if err != nil {
		log.Fatal(err.Error())
	}
}
