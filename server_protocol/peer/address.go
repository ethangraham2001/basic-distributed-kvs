package peer

import (
	"fmt"
	"net"
)

// APIEndpoint is the only available Endpoint for now.
const APIEndpoint string = "/api/"

// Address wraps an IP-address and Port pair.
type Address struct {
	IP   net.IP // IP can be ipv4 or ipv6
	Port uint16
}

func (addr *Address) String() string {
	return fmt.Sprintf("http://%s:%d", addr.IP.String(), addr.Port)
}
