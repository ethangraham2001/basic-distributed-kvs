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

func handlePut(c *client.Client, key string, filePath string) {
	data, err := util.ReadFile(filePath)
	if err != nil {
		log.Fatal(err.Error())
	}

	keyHash := util.HashKey(key)
	peerID := keyHash % c.NumPeers

	err = c.MakePutRequest(peerID, key, data)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func handleGet(c *client.Client, key string, filePath string) {
	keyHash := util.HashKey(key)
	peerID := keyHash % c.NumPeers

	data, err := c.MakeGetRequest(peerID, key)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = util.WriteFile(filePath, data)
	if err != nil {
		log.Fatal(err.Error())
	}
}
