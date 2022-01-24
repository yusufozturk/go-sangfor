package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/hex"
	"errors"
	"math/big"
	"strings"
)

// Authenticate Function to make authentication request
func (client *Client) Authenticate(username, encryptedPassword string) error {
	// Get Auth Body
	authRequest := SangforAuthReq{}
	authRequest.Auth.PasswordCredentials.Username = username
	authRequest.Auth.PasswordCredentials.Password = encryptedPassword

	// Make Authentication
	authResponse := &SangforAuthResponse{}
	_, err := client.Client.R().
		SetResult(authResponse).
		SetBody(authRequest).
		SetHeader("Content-Type", "application/json").
		SetHeader("Cookie", "aCMPAuthToken="+getUUID()).
		Post(client.BaseAPIURL + "/authenticate")
	if err != nil {
		return err
	}

	// Check Response
	if authResponse == nil {
		return errors.New("no response")
	}

	// Check Token
	if authResponse.Data.Access.Token.ID == "" {
		return errors.New("token is not available")
	}

	// Set Token
	client.Token = authResponse.Data.Access.Token.ID

	return nil
}

// GetPublicKey Gets Sangfor API Public Key
func GetPublicKey(client *Client) (string, error) {
	// Get Public Key
	sangforPK := &SangforPK{}
	_, err := client.Client.R().
		SetResult(sangforPK).
		Get(client.BaseAPIURL + "/public-key")
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
