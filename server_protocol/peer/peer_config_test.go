package peer

import (
	"net"
	"testing"
)

func TestReadConfig(t *testing.T) {
	data := []byte(`
    Id: 0
    Peers:
        1: 
            IP: 127.0.0.1
            Port: 8000
        2: 
            IP: 127.0.0.1
            Port: 12345
    `)

	expected := newPeer[int, int](0)
	expected.addConnection(1, Address{IP: net.ParseIP("127.0.0.1"), Port: 8000})
	expected.addConnection(2, Address{IP: net.ParseIP("127.0.0.1"), Port: 12345})
	actual, err := readConfig[int, int](data)

	if err != nil {
		t.Fatalf("Parsing valid yaml should not return error")
	}

	if actual.ID != expected.ID {
		t.Fatalf("IDs should be equal. Expected %d, got %d", actual.ID, expected.ID)
	}

	if expected.connections == nil {
		t.Fatalf("expected connections should not be nil")
	}

	if actual.connections == nil {
		t.Fatalf("actual connections should not be nil")
	}

	for id, addr := range actual.connections {
		if !addr.IP.Equal(expected.connections[id].IP) {
			t.Fatalf("IP address for id=%d does not match; ", id)
		}

		if addr.Port != expected.connections[id].Port {
			t.Fatalf("Port for id=%d does not match; ", id)
		}
	}
}
