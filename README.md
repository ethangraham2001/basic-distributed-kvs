# Basic Distributed Key-Value Store

Currently working on implementing the server-side protocol that will serve
client requests.

## Request API

### `Get()` API

Expects a `GET` http-request to `/api/<key>` and returns binary data in the
form of an http response with `Content-Type: application/octet-stream` in the
header.

### `Put()` API

Expects a `PUT` http-request to `/api/<key>` with binary data in the form of
an http request containing `Content-Type: application/octet-stream` in the 
header.

