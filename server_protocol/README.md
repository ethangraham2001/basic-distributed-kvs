# Configuration

Currently configuration is done via YAML - the contents should be placed in
`config.yaml` in the top level of the server protocol directory. Here is an
example

```yaml
# Example configuration for a four-peer system. 
Id: 0
Peers:
  0:
    IP: 127.0.0.1
    Port: 50000
  1:
    IP: 127.0.0.1
    Port: 50001
  2:
    IP: 127.0.0.1
    Port: 50002
  3:
    IP: 127.0.0.1
    Port: 50003
```

The `Id` field should be different on every system, and the peer should appear
in its own `Peers` map.

