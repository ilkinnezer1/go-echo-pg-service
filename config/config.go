package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// handleError simply output the error message
func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GetCredentials gives credentials for the login process
func GetCredentials() (string, string) {
	err := godotenv.Load()
	handleError(err)
	var username = os.Getenv("APP_USERNAME")
	var password = os.Getenv("APP_PASSWD")

	return username, password
}
