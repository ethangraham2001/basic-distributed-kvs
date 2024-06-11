package datastore

import (
	"testing"
)

func TestAddElement(t *testing.T) {
	datastore := NewInMemoryDataStore[int64, string]()

	datastore.Put(0, "hello")
	datastore.Put(1, "world")
	datastore.Put(1, "!")

	expected := "hello"
	actual, err := datastore.Get(0)
	if err != nil {
		t.Fatalf("Get(0) should not return an error")
	}
	if actual != expected {
		t.Fatalf("Get(0): expected %s, got %s", expected, actual)
	}

	expected = "!"
	actual, err = datastore.Get(1)
	if err != nil {
		t.Fatalf("Get(1) should not return an error")
	}
	if actual != expected {
		t.Fatalf("Get(1): expected %s, got %s", expected, actual)
	}
}
