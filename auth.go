package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"math/big"
	"net/http"
	"strings"
)

// Authenticate Function to make authentication request
func (client *Client) Authenticate(username, encryptedPassword string) error {
	// Get Auth Body
	authRequest := SangforAuthReq{}
	authRequest.Auth.PasswordCredentials.Username = username
	authRequest.Auth.PasswordCredentials.Password = encryptedPassword

	// Convert struct to json data
	reqBody, err := json.Marshal(authRequest)

	if err != nil {
		return err
	}

	// Create request for authentication
	req, err := http.NewRequest("POST", client.BaseAPIURL+"/authenticate", bytes.NewReader(reqBody))
	if err != nil {
		return err
	}

	// Set request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", "aCMPAuthToken="+getUUID())

	defer req.Body.Close()

	// Make authentication
	resp, err := client.Client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Convert JSON response to struct
	authResponse := &SangforAuthResponse{}
	err = json.Unmarshal(respBody, authResponse)
	if err != nil {
		return err
	}

	// Check Response
	if authResponse == nil {
		return errors.New("no response")
	}

	// Check Token
	if authResponse.Data.Access.Token.ID != "" {
		// Set Token
		client.Token = authResponse.Data.Access.Token.ID
		return nil
	}

	return errors.New("token is not available")
}

// GetPublicKey Gets Sangfor API Public Key
func GetPublicKey(client *Client) (string, error) {
	// Get Public Key
	resp, err := client.Client.Get(client.BaseAPIURL + "/public-key")
	if err != nil {
		return "", err
	}

	// Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Convert JSON response to struct
	sangforPK := &SangforPK{}
	err = json.Unmarshal(respBody, sangforPK)
	if err != nil {
		return "", err
	}

	return sangforPK.Data.PublicKey, nil
}

// GetEncryptedPassword Gets Sangfor Encrypted Password
func GetEncryptedPassword(password, publicKey string) (string, error) {
	// Get Properties
	N := strings.TrimSuffix(publicKey, "\n")
	E := 65537

	// Convert Public Key
	bigN := new(big.Int)
	_, ok := bigN.SetString(N, 16)
	if !ok {
		return "", errors.New("can not convert public key")
	}

	// Get Public Key
	pub := rsa.PublicKey{
		N: bigN,
		E: E,
	}

	// Encrypt Password
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, &pub, []byte(password))
	if err != nil {
		return "", nil
	}

	// Encode Encrypted Password
	return hex.EncodeToString(cipherText), nil
}

// SangforPK represents Sangfor Public Key Format
type SangforPK struct {
	Data struct {
		PublicKey string `json:"public_key"`
	} `json:"data"`
}
