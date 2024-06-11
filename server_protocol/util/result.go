// Package util contains some simple utilitary functions definitions.
package util

// Result type
type Result[T any] struct {
	Value T
	Err   error
}
