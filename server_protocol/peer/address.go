package peer

import (
	"net"
)

// Address wraps an IP-address and Port pair. The IP-address field can
// contain either an ipv4 or ipv6 address.
type Address struct {
	IP   net.IP
	Port uint16
}
