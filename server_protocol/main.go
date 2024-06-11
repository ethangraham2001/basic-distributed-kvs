// Package main is the main package of the project.
package main

import (
	"fmt"
    "github.com/ethangraham2001/distributed_kvs_server_protocol/peer"
)

const filepath string = "config.yaml"

func main() {
    p, err := peer.ReadConfigFromFile[uint64, uint64](filepath)

    if err != nil {
        panic("failed to initialize peer!")
    }

    fmt.Println(p.ID)
}
