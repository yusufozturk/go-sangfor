package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define Flags
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	hostFlag := flag.String("host", "", "SCP Cloud Platform Address")
	usernameFlag := flag.String("username", "", "SCP Cloud Platform Username")
	passwordFlag := flag.String("password", "", "SCP Cloud Platform Password")

	// Parse Flags
	flag.Parse()

	// Check Host
	if *hostFlag == "" {
		fmt.Println(errors.New("host parameter is missing"))
		return
	}

	// Check Username
	if *usernameFlag == "" {
		fmt.Println(errors.New("username parameter is missing"))
		return
	}

	// Check Password
	if *passwordFlag == "" {
		fmt.Println(errors.New("password parameter is missing"))
		return
	}

	// Get Client
	client := GetAPIClient(*hostFlag)

	// Get Authentication Token
	err := run(client, *usernameFlag, *passwordFlag)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func run(client *Client, username, password string) error {
	// Get Public Key for Authentication
	publicKey, err := GetPublicKey(client)
	if err != nil {
		return err
	}

	// Encrypt Password
	encryptedPassword, err := GetEncryptedPassword(password, publicKey)
	if err != nil {
		return err
	}

	// Authenticate
	err = client.Authenticate(username, encryptedPassword)
	if err != nil {
		return err
	}

	// Output
	fmt.Println("Encrypted Password: ", encryptedPassword)
	fmt.Println("---------------------------------------")
	fmt.Println("Token: ", client.Token)

	return nil
}
