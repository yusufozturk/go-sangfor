package main

import (
	"crypto/tls"
	"net/http"
	"time"
)

// Client represents a connection to Sangfor
type Client struct {
	Client     *http.Client
	BaseAPIURL string
	PublicKey  string
	Token      string
	Valid      bool
}

// GetAPIClient Gets Sangfor API Client
func GetAPIClient(host string) *Client {
	// Create custom transport layer configuration
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
			MinVersion:         getMinTLSVersion(),
		},
	}

	// Create HTTP Client
	httpClient := &http.Client{
		Transport: tr,
	}

	// Sangfor API Client
	client := &Client{
		Client: httpClient,
		Valid:  true,
	}

	// Set Base API
	client.BaseAPIURL = host + "/janus"

	return client
}

// getMinTLSVersion Gets Minimum TLS Version
func getMinTLSVersion() uint16 {
	return tls.VersionTLS12
}

// SangforAuthResponse represents the auth response
type SangforAuthResponse struct {
	Message string `json:"message"`
	Data    struct {
		Access struct {
			Token struct {
				IssuedAt time.Time `json:"issued_at"`
				Expires  time.Time `json:"expires"`
				ID       string    `json:"id"`
				Tenant   struct {
					ID          string `json:"id"`
					Description string `json:"description"`
					Name        string `json:"name"`
					Enabled     bool   `json:"enabled"`
				} `json:"tenant"`
				AuditIds []string `json:"audit_ids"`
			} `json:"token"`
			User struct {
				Username   string        `json:"username"`
				Name       string        `json:"name"`
				ID         string        `json:"id"`
				RolesLinks []interface{} `json:"roles_links"`
				Roles      []struct {
					Name string `json:"name"`
				} `json:"roles"`
			} `json:"user"`
		} `json:"access"`
	} `json:"data"`
	Code int `json:"code"`
}

// SangforAuthReq represents the auth request
type SangforAuthReq struct {
	Auth struct {
		PasswordCredentials struct {
			Username string `json:"username"`
			Password string `json:"password"`
		} `json:"passwordCredentials"`
	} `json:"auth"`
}
