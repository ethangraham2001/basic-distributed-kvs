// Package util provides utilitary functions
package util

import (
	"io"
	"os"
)

// ReadFile reads a file from persistent storage.
func ReadFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)

	if err != nil {
		return []byte{}, err
	}
	defer file.Close()

	contents, err := io.ReadAll(file)
	if err != nil {
		return []byte{}, err
	}

	return contents, nil
}
