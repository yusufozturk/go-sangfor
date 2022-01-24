package main

import (
	"testing"
)

func TestGetUUID(t *testing.T) {
	// Get UUID
	uuid := getUUID()
	if uuid == "" {
		t.Fatal()
	}
}
