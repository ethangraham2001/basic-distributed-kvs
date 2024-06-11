// Package address contains definition of address type representing an
// (IP, Port) pair
package address

import (
	"fmt"
	"net"
)

// Address wraps an IP-address and Port pair.
type Address struct {
	IP   net.IP // IP can be ipv4 or ipv6
	Port uint16
}

// NewAddress initializes an Address and returns it
func NewAddress(ipStr string, port uint16) Address {
    return Address{ IP: net.ParseIP(ipStr), Port: port }
}

// String returns the string representation of an AddressTpe.
// Of the form "<IP>:<Port>"
func (addr *Address) String() string {
    return fmt.Sprintf("http://%s:%d", addr.IP.String(), addr.Port)
}
