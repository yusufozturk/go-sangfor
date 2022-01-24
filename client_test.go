package main

import (
	"crypto/tls"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIClient(t *testing.T) {
	// Get Client
	client := GetAPIClient("host")

	expectedClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
				MinVersion:         getMinTLSVersion(),
			},
		},
	}

	// Check Client
	if client.BaseAPIURL != "host/janus" {
		t.Fatal()
	}

	if !reflect.DeepEqual(client.Client, expectedClient) || !client.Valid {
		t.Fatal()
	}
}
