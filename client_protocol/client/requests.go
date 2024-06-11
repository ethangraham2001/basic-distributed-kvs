package client

type getRequest[K comparable] struct {
	Key K `json:"key"`
}

type putRequest[K comparable] struct {
	Key K `json:"key"`
}
