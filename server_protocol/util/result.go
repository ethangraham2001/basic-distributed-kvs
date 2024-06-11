// Package util contains some simple utilitary functions definitions.
package util

// Result type that contains either a value Value field, or a non-nil Err
type Result[T any] struct {
	Value T
	Err   error
}
