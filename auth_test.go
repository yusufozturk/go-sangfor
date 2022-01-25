package main

import (
	"testing"
)

func TestAuthenticate_OK(t *testing.T) {
	// Test Response
	authResponse := &SangforAuthResponse{}
	authResponse.Data.Access.Token.ID = "Test"

	// Mock Database
	db := MockDatabase{}
	db.AddToDatabase("", authResponse)

	// Get Client
	client := GetAPIClient("host")
	client.Client.Transport = NewMockTransport(200, map[string]string{}, db.Database)

	// Authenticate
	err := client.Authenticate("", "")

	// Check Errors
	if err != nil {
		t.Fatal(err)
		return
	}
}

func TestAuthenticate_NoToken(t *testing.T) {
	// Test Response
	authResponse := &SangforAuthResponse{}
	authResponse.Data.Access.Token.ID = ""

	// Mock Database
	db := MockDatabase{}
	db.AddToDatabase("", authResponse)

	// Get Client
	client := GetAPIClient("host")
	client.Client.Transport = NewMockTransport(200, map[string]string{}, db.Database)

	// Authenticate
	err := client.Authenticate("", "")

	// Check Errors
	if err == nil {
		t.Fatal(err)
	}
}

func TestGetPublicKey_OK(t *testing.T) {
	// Test Response
	sangforPK := SangforPK{}
	sangforPK.Data.PublicKey = "PublicKey"

	// Mock Database
	db := MockDatabase{}
	db.AddToDatabase("", sangforPK)

	// Get Client
	client := GetAPIClient("host")
	client.Client.Transport = NewMockTransport(200, map[string]string{}, db.Database)

	// Get Public Key
	publicKey, err := GetPublicKey(client)

	// Check Public Key
	if publicKey == "PublicKey" {
		return
	}

	// Check Errors
	if err == nil {
		t.Fatal("connection is failed")
		return
	}

	t.Fatal(err)
}

func TestGetPublicKey_Error(t *testing.T) {
	// Get Client
	client := GetAPIClient("host")
	client.Client.Transport = NewMockTransport(500, map[string]string{}, map[string][]byte{})

	// Get Public Key
	_, err := GetPublicKey(client)

	// Check Errors
	if err == nil {
		t.Fatal("connection is failed")
		return
	}
}

func TestGetPublicKey_JSONError(t *testing.T) {
	// Test Response
	sangforPK := SangforPK{}
	sangforPK.Data.PublicKey = "PublicKey"

	// Get Client
	client := GetAPIClient("host")
	client.Client.Transport = NewMockTransport(200, map[string]string{}, map[string][]byte{})

	// Get Public Key
	_, err := GetPublicKey(client)

	// Check Errors
	if err == nil {
		t.Fatal("json validation is failed")
		return
	}
}

func TestGetEncryptedPassword(t *testing.T) {
	// Public Key
	publicKey := "D9905621B5800B5BEC903FB5E96C42DF23B6B0ABC9878C7310A62254DD0F8B54C6027C5A0C0511" +
		"8056E4EE72DFFD3DD9CF61EC83AD73BD2A8988872FC6015F0645F5A8E71A5F8A475C6F46F98F653250C51CE" +
		"BB88534BA89C2C75FD29F7866618A1D3FEC53EEC4C02FB33D8CBD6585FAD36E49546E4E984FB867A755ADA0" +
		"89386BF94E748EF996018DC8F47C321C452FE57FB36B3D2F9635CA94436D422DD2746DE0B67EF420A77DBAF" +
		"B110DD03989A712A266841360EDA612CDA0E435C649116B3988CE27EADC69ED1008FBF91F0B31B903ED8151" +
		"11C8540CF72C58FC601F72C3A019448399AB0DB0AC83AEC11753E06AF0E0CDE46388F07C47BD74F0A75095"

	// Get Password
	encpass, err := GetEncryptedPassword("password", publicKey)

	// Check Password
	if encpass != "" {
		return
	}

	// Check Errors
	if err == nil {
		t.Fatal("encryption is failed")
		return
	}

	t.Fatal(err)
}
