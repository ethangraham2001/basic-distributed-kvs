# Client Program

## Running the program

### Configuration

The client program's peer-list is configured by reading from `config.yaml`
placed in the same directory as the binary. It expects the following format

```yaml
# Example configuration for a four-peer system. 
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

### Usage
```bash
# Gets a blob from the key-value store and writes it to an output file
$ ./distributed_kvs_client_protocol GET <key> <output_filename>

# reads a blob from an input file and puts it in the key-value store
$ ./distributed_kvs_client_protocol PUT <key> <input_filename>
```

## Functionality

The client program performs an MD5 hash on the key and computes its modulus
with respect to the number of peers in the key value store to decide where it
should be stored.

## TODO

- [ ] Handle failure: when the peer is offline, the program will fail. Once
replication is implemented for the key-value store, a priority list will be 
used.

