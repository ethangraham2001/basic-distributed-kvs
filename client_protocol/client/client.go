// Package client contains definition of Client struct
package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

    "github.com/ethangraham2001/distributed_kvs_client_protocol/address"
)

const apiEndpoint string = "/api/"

const contentType string = "Content-Type"
const contentTypeJSON string = "application/json"
const contentTypeOctetStream string = "application/octet-stream"

// Client wraps all data needed for communication with server.
// K generic type is the type of the key
type Client[K comparable] struct {
	// connections maps Peer identifiers to their Address
	connections map[uint32]address.Address
}

// NewClient initializes and returns a new Client
func NewClient[K comparable]() Client[K] {
	return Client[K]{connections: make(map[uint32]address.Address)}
}

// MakeGetRequest makes a get request to Peer with ID `peerId` and
// handles the response.
// Returns a byte array on success, or an error on failure.
// Does not attempt to find the correct Peer based on hashing algorithm,
// assuming that this is done by the caller.
func (c *Client[K]) MakeGetRequest(peerID uint32, key K) ([]byte, error) {
	addr, valid := c.connections[peerID]
	if !valid {
		return []byte{}, errors.New("Peer not found")
	}

	getReq := getRequest[K]{Key: key}
	jsonData, err := json.Marshal(getReq)
	if err != nil {
		return []byte{}, nil
	}

	req, err := http.NewRequest(http.MethodGet, addr.String()+apiEndpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return []byte{}, err
	}

	req.Header.Set(contentType, contentTypeJSON)

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
func (c *Client[K]) MakePutRequest(peerID uint32, key K, value []byte) error {
	addr, valid := c.connections[peerID]

	if !valid {
		return errors.New("Peer not found")
	}

	var requestBody bytes.Buffer

	writer := multipart.NewWriter(&requestBody)

	putReq := putRequest[K]{Key: key}

	jsonPart, err := writer.CreatePart(map[string][]string{contentType: {contentTypeJSON}})
	if err != nil {
		return err
	}

	jsonEncoder := json.NewEncoder(jsonPart)
	if err := jsonEncoder.Encode(putReq); err != nil {
		return err
	}

	octetStreamPart, err := writer.CreatePart(map[string][]string{contentType: {contentTypeOctetStream}})
	if err != nil {
		return err
	}
	// write the raw byte array into the octet-stream section of the request
	_, err = octetStreamPart.Write(value)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, addr.String()+apiEndpoint, &requestBody)
	if err != nil {
		return err
	}

	req.Header.Set(contentType, writer.FormDataContentType())

	client := &http.Client{}

	resp, err := client.Do(req)
    if err != nil {
        return err
    }

	if resp.StatusCode != http.StatusOK {
		errorMsg := fmt.Sprintf("PUT request failed. Code = %s", resp.Status)
		return errors.New(errorMsg)
	}

	return nil
}

func (c *Client[K]) AddConnection(peerID uint32, peerAddr address.Address) {
	c.connections[peerID] = peerAddr
}
