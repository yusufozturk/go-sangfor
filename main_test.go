package main

import (
	"testing"
)

func TestRun_OK(t *testing.T) {
	// Test Response: Public Key
	sangforPK := SangforPK{}
	sangforPK.Data.PublicKey = "D9905621B5800B5BEC903FB5E96C42DF23B6B0ABC9878C7310A62254DD0F8B54C6027C5A0C0511" +
		"8056E4EE72DFFD3DD9CF61EC83AD73BD2A8988872FC6015F0645F5A8E71A5F8A475C6F46F98F653250C51CE" +
		"BB88534BA89C2C75FD29F7866618A1D3FEC53EEC4C02FB33D8CBD6585FAD36E49546E4E984FB867A755ADA0" +
		"89386BF94E748EF996018DC8F47C321C452FE57FB36B3D2F9635CA94436D422DD2746DE0B67EF420A77DBAF" +
		"B110DD03989A712A266841360EDA612CDA0E435C649116B3988CE27EADC69ED1008FBF91F0B31B903ED8151" +
		"11C8540CF72C58FC601F72C3A019448399AB0DB0AC83AEC11753E06AF0E0CDE46388F07C47BD74F0A75095"

	// Test Response: Auth Response
	authResponse := &SangforAuthResponse{}
	authResponse.Data.Access.Token.ID = "Test"

	// Mock Database
	db := MockDatabase{}
	db.AddToDatabase("host/janus/public-key", sangforPK)
	db.AddToDatabase("host/janus/authenticate", authResponse)

	// Get Client
	client := GetAPIClient("host")
	client.Client.Transport = NewMockTransport(200, map[string]string{}, db.Database)

	// Test Run
	err := run(client, "", "")
	if err != nil {
		t.Fatal(err)
	}
}

func TestRun_Error(t *testing.T) {
	// Test Response: Public Key
	sangforPK := SangforPK{}
	sangforPK.Data.PublicKey = "Test"

	// Test Response: Auth Response
	authResponse := &SangforAuthResponse{}
	authResponse.Data.Access.Token.ID = "Test"

	// Mock Database
	db := MockDatabase{}
	db.AddToDatabase("host/janus/public-key", sangforPK)
	db.AddToDatabase("host/janus/authenticate", authResponse)

	// Get Client
	client := GetAPIClient("host")
	client.Client.Transport = NewMockTransport(200, map[string]string{}, db.Database)

	// Test Run
	err := run(client, "", "")
	if err == nil {
		t.Fatal("error validation is failed")
	}
}

func TestRun_PK_Error(t *testing.T) {
	// Test Response: Auth Response
	authResponse := &SangforAuthResponse{}
	authResponse.Data.Access.Token.ID = "Test"

	// Mock Database
	db := MockDatabase{}
	db.AddToDatabase("host/janus/public-key", "Test")
	db.AddToDatabase("host/janus/authenticate", authResponse)

	// Get Client
	client := GetAPIClient("host")
	client.Client.Transport = NewMockTransport(200, map[string]string{}, db.Database)

	// Test Run
	err := run(client, "", "")
	if err == nil {
		t.Fatal("error validation is failed")
	}
}

func TestRun_Response_Error(t *testing.T) {
	// Test Response: Public Key
	sangforPK := SangforPK{}
	sangforPK.Data.PublicKey = "D9905621B5800B5BEC903FB5E96C42DF23B6B0ABC9878C7310A62254DD0F8B54C6027C5A0C0511" +
		"8056E4EE72DFFD3DD9CF61EC83AD73BD2A8988872FC6015F0645F5A8E71A5F8A475C6F46F98F653250C51CE" +
		"BB88534BA89C2C75FD29F7866618A1D3FEC53EEC4C02FB33D8CBD6585FAD36E49546E4E984FB867A755ADA0" +
		"89386BF94E748EF996018DC8F47C321C452FE57FB36B3D2F9635CA94436D422DD2746DE0B67EF420A77DBAF" +
		"B110DD03989A712A266841360EDA612CDA0E435C649116B3988CE27EADC69ED1008FBF91F0B31B903ED8151" +
		"11C8540CF72C58FC601F72C3A019448399AB0DB0AC83AEC11753E06AF0E0CDE46388F07C47BD74F0A75095"

	// Mock Database
	db := MockDatabase{}
	db.AddToDatabase("host/janus/public-key", sangforPK)
	db.AddToDatabase("host/janus/authenticate", "Test")

	// Get Client
	client := GetAPIClient("host")
	client.Client.Transport = NewMockTransport(200, map[string]string{}, db.Database)

	// Test Run
	err := run(client, "", "")
	if err == nil {
		t.Fatal("error validation is failed")
	}
}
