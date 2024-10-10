package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func generateAccessToken(userID string, clientSecret string) string {
	salted := clientSecret + userID

	hasher := md5.New()
	hasher.Write([]byte(salted))
	hashBytes := hasher.Sum(nil)

	hashString := hex.EncodeToString(hashBytes)

	accessToken := userID + "_" + hashString

	return accessToken
}

func main() {
	// Sample clientSecret
	clientSecret := "ETT_SECRET_SALT_"

	// Sample userIDs
	userIDs := []string{
		"12345",
		"67890",
		"54321",
		"98765",
		"11223",
	}

	// Generate and print access tokens
	for _, userID := range userIDs {
		token := generateAccessToken(userID, clientSecret)
		fmt.Println(token)
	}
}
