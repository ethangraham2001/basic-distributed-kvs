# Basic Distributed Key-Value Store

Very simple implementation of a distributed key-value store leveraging 
consistent hashing for mapping a key to a peer.

Every key has a preference list of $N$ peers, which is the range
`[hash(key) % num_peers, hash(key) + N % num_peers]`.
The hash function used is MD5. 
- If a peer is not the last in this range, it will attempt to replicate the
data to the next peer.
- If a client's request to a peer fails, it will attempt the next peer in the
preference list either until success or until there are no peers left.

## Demo

![Demo Video](res/demo.gif)

### Explanation 

- Fire up two peers `Peer_0` and `Peer_1`
- Client makes a `Put()` call for a blob of data with a key that maps to 
`Peer_0`. This request is only made to `Peer_0`
- `Peer_0` replicates the data via. put request to `Peer_1` (if the preference
list were longer, `Peet_1` would do the same)
- `Peer_0` is shut down
- Client makes a `Get()` call for the same key. Its request for `Peer_0` fails,
so it makes a successful call to `Peer_1` which is the next in the preference
list for that key
- The resulting data is written to a new file

## Request API

All communication happens over http

### `Get()` API

Expects a `GET` http-request to `/api/<key>` and returns binary data in the
form of an http response with `Content-Type: application/octet-stream` in the
header.

### `Put()` API

Expects a `PUT` http-request to `/api/<key>` with binary data in the form of
an http request containing `Content-Type: application/octet-stream` in the 
header.

## Running the program

Check the `/client_protocol` and `/server_protocol` readme files for 
information on configuration.

