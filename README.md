# Basic Distributed Key-Value Store

Currently working on implementing the server-side protocol that will serve
client requests.

## Request API

Client-to-peer and peer-to-peer request are JSON formatted

### Get() API

Expects an http `GET` with json request body

```json
{
    key: "<Some Key>"
}
```

Returns binary data.

### Put() API

Expects a http `PUT` with json request body

```json
{
    key: "<Some Key>"
}
```

with binary data in the request.

