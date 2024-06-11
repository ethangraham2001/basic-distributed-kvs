package peer

import (
	"io"
	"net"
	"os"

	"gopkg.in/yaml.v3"
)

// Config represents the structure of the config.yaml file used for
// configuring peers
type Config struct {
	ID    uint32
	Peers map[uint32]marshalledAddress
}

// same structure as address, but taking a string as parameter for
// compatibility with the yaml package. The IP address is parsed into
// a net.IP object manually
type marshalledAddress struct {
	IP   string
	Port uint32
}

// ReadConfigFromFile reads a .yaml configuration file and returns the
// resulting Peer object or an error.
func ReadConfigFromFile[K comparable, V any](filepath string) (Peer[K, V], error) {
	var p Peer[K, V]

	file, err := os.Open(filepath)
	if err != nil {
		return p, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	return readConfig[K, V](data)
}

// readConfig returns a Peer object parsed from a raw .yaml file or returns
// an error
func readConfig[K comparable, V any](data []byte) (Peer[K, V], error) {
	c := Config{}
	err := yaml.Unmarshal(data, &c)

	var p Peer[K, V]
	if err != nil {
		return p, err
	}

	p = newPeer[K, V](c.ID)
	for id, marshalledAddress := range c.Peers {
		ip := net.ParseIP(marshalledAddress.IP)
		addr := Address{IP: ip, Port: uint16(marshalledAddress.Port)}
		p.addConnection(id, addr)
	}

	return p, nil
}
