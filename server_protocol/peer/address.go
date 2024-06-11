package peer

import (
	"net"
)

// Address wraps an IP-address and Port pair.
type Address struct {
	IP   net.IP // IP can be ipv4 or ipv6
	Port uint16
}
