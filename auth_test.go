package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"reflect"
	"testing"
)

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

	if encpass == "" {
		if err.Error() == "can not convert public key" {
			t.Fatal(err)
		}
		if err == nil {
			t.Fatal("encrypt fail")
		}
	}
}

func TestPrepareAuthRequest(t *testing.T) {
	expectedReqBody := SangforAuthReq{}
	expectedReqBody.Auth.PasswordCredentials.Username = "test"
	expectedReqBody.Auth.PasswordCredentials.Password = "test123"

	// Convert struct to json data
	expectedReqBodyJson, err := json.Marshal(expectedReqBody)
	if err != nil {
		t.Fatal(err)
	}

	expectedReq, err := http.NewRequest("POST", "test.com/api/authenticate", bytes.NewReader(expectedReqBodyJson))
	if err != nil {
		t.Fatal(err)
	}

	expectedReq.Body.Close()

	req, err := prepareAuthRequest("test", "test123", "test.com/api")
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(req.Body, expectedReq.Body) || req.Method != expectedReq.Method || req.URL.String() != expectedReq.URL.String() {
		t.Fatal(err)
	}
}
