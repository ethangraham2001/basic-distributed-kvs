package client

import (
	"io"
	"net"
	"os"

	"github.com/ethangraham2001/distributed_kvs_client_protocol/address"
	"gopkg.in/yaml.v3"
)

// Config represents the structure of the config.yaml file used for
// configuring peers
type Config struct {
	Peers map[uint32]marshalledAddress `yaml:"Peers"`
}

// same structure as address, but taking a string as parameter for
// compatibility with the yaml package. The IP address is parsed into
// a net.IP object manually
type marshalledAddress struct {
	IP   string `yaml:"IP"`
	Port uint32 `yaml:"Port"`
}

// ReadConfigFromFile reads a .yaml configuration file and returns the
// resulting Client struct or an error.
func ReadConfigFromFile(filepath string) (Client, error) {
	var c Client

	file, err := os.Open(filepath)
	if err != nil {
		return c, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	return readConfig(data)
}

// readConfig returns a Client struct parsed from a raw .yaml file or returns
// an error
func readConfig(data []byte) (Client, error) {
	config := Config{}
	err := yaml.Unmarshal(data, &config)

	var c Client
	if err != nil {
		return c, err
	}

	c = NewClient(uint32(len(config.Peers)))
	for id, marshalledAddress := range config.Peers {
		ip := net.ParseIP(marshalledAddress.IP)
		addr := address.Address{IP: ip, Port: uint16(marshalledAddress.Port)}
		c.addConnection(id, addr)
	}
	c.NumPeers = uint32(len(config.Peers))
	return c, nil
}
