package main

import (
	"crypto/tls"
	"time"

	"github.com/go-resty/resty/v2"
)

// GetAPIClient Gets Sangfor API Client
func GetAPIClient(host string) *Client {
	// Resty Client
	restyClient := resty.New()
	restyClient.SetTLSClientConfig(
		&tls.Config{
			InsecureSkipVerify: true,
			MinVersion:         GetMinTLSVersion(),
		})

	// Sangfor API Client
	client := &Client{
		Client: restyClient,
		Valid:  true,
	}

	// Set Base API
	client.BaseAPIURL = host + "/janus"

	return client
}

// GetMinTLSVersion Gets Minimum TLS Version
func GetMinTLSVersion() uint16 {
	return tls.VersionTLS12
}

// Client represents a connection to Sangfor
type Client struct {
	Client     *resty.Client
	BaseAPIURL string
	PublicKey  string
	Token      string
	Valid      bool
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
