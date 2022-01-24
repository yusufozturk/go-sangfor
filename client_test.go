package main

import (
	"testing"
)

func TestGetAPIClient(t *testing.T) {
	// Get Client
	client := GetAPIClient("host")

	// Check Client
	if client.BaseAPIURL != "host/janus" {
		t.Fatal()
	}
}
