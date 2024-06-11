// Package client contains definition of Client struct
package client

import (
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/ethangraham2001/distributed_kvs_client_protocol/address"
)

const apiEndpoint string = "/api/"

const contentType string = "Content-Type"
const contentTypeJSON string = "application/json"
const contentTypeOctetStream string = "application/octet-stream"

// Client wraps all data needed for communication with server.
// K generic type is the type of the key
type Client struct {
	// connections maps Peer identifiers to their Address
	connections map[uint32]address.Address
}

// NewClient initializes and returns a new Client
func NewClient() Client {
	return Client{connections: make(map[uint32]address.Address)}
}

// MakeGetRequest makes a get request to Peer with ID `peerId` and
// handles the response.
// Returns a byte array on success, or an error on failure.
// Does not attempt to find the correct Peer based on hashing algorithm,
// assuming that this is done by the caller.
func (c *Client) MakeGetRequest(peerID uint32, key string) ([]byte, error) {
	addr, valid := c.connections[peerID]
	if !valid {
		return []byte{}, errors.New("Peer not found")
	}

    path := apiEndpoint + key
    requestBody := []byte{}
	req, err := http.NewRequest(http.MethodGet, addr.String() + path, bytes.NewBuffer(requestBody))
	if err != nil {
		return []byte{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}

// MakePutRequest makes a PUT request for (key, value) in `peerId`th Peer
// Does not attempt to find the correct Peer based on hashing algorithm,
// assuming that this is done by the caller.
func (c *Client) MakePutRequest(peerID uint32, key string, value []byte) error {
	addr, valid := c.connections[peerID]

	if !valid {
		return errors.New("Peer not found")
	}

    path := apiEndpoint + key
    
    req, err := http.NewRequest(http.MethodPut, addr.String() + path, bytes.NewBuffer(value))
    if err != nil {
        log.Printf("Failed to create request. %s", err.Error())
        return err
    }

    req.Header.Set(contentType, contentTypeOctetStream)
    client := &http.Client{}
    
    resp, err := client.Do(req)
    if err != nil {
        log.Printf("Failed to communicate with server. %s", err.Error())
        return err
    }
    defer resp.Body.Close()

    _, err = io.ReadAll(resp.Body)
    if err != nil {
        log.Printf("failed to read response body %s", err.Error())
        return err
    }

    return nil
}

// AddConnection adds a (peerID -> peerAddress) mapping to the client.
func (c *Client) AddConnection(peerID uint32, peerAddr address.Address) {
	c.connections[peerID] = peerAddr
}
