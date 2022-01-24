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
		panic(errors.New("host parameter is missing"))
	}

	// Check Username
	if *usernameFlag == "" {
		panic(errors.New("username parameter is missing"))
	}

	// Check Password
	if *passwordFlag == "" {
		panic(errors.New("password parameter is missing"))
	}

	// Get Authentication Token
	run(*hostFlag, *usernameFlag, *passwordFlag)
}

func run(host, username, password string) {
	// Get Client
	client := GetAPIClient(host)

	// Get Public Key for Authentication
	publicKey, err := GetPublicKey(client)
	if err != nil {
		panic(err)
	}

	// Encrypt Password
	encryptedPassword, err := GetEncryptedPassword(password, publicKey)
	if err != nil {
		panic(err)
	}

	// Authenticate
	err = client.Authenticate(username, encryptedPassword)
	if err != nil {
		panic(err)
	}

	// Output
	fmt.Println(encryptedPassword)
	fmt.Println(client.Token)
}
