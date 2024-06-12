// Package util provides utilitary functions
package util

import (
	"io"
	"log"
	"os"
)

// ReadFile reads a file from persistent storage.
func ReadFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)

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

// WriteFile writes a file to persistent storage.
func WriteFile(filePath string, data []byte) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		log.Printf("Unable to open %s", filePath)
		return err
	}

	defer func() {
		log.Printf("Closing %s", filePath)
		file.Close()
	}()

	_, err = file.Write(data)
	if err != nil {
		log.Printf("Failure writing to %s", filePath)
		return err
	}

	return nil
}
