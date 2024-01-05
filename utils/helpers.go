package utils

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func ValidateUsername(username string) error {
	// Check for special characters in the username
	if strings.ContainsAny(username, "!@#$%^&*()_+-=[]{}|;:'\",.<>?/") {
		return fmt.Errorf("Username cannot contain special characters")
	}
	return nil
}

func ValidatePassword(password string) error {
	// Check for special characters in the password
	if strings.ContainsAny(password, "!@#$%^&*()_+-=[]{}|;:'\",.<>?/") {
		return fmt.Errorf("Password cannot contain special characters")
	}
	return nil
}

func ValidatePhoneNumber(phoneNumber string) error {
	// Simple phone number validation using a regular expression
	// You can replace this with a more sophisticated validation method
	// or use a third-party library for phone number validation.
	phoneRegex := `^\+[1-9]\d{1,14}$`
	match, err := regexp.MatchString(phoneRegex, phoneNumber)
	if err != nil {
		return err
	}
	if !match {
		return fmt.Errorf("Invalid phone number format")
	}
	return nil
}

func Getenv(key string) []byte {
	item := []byte(os.Getenv(key))

	// Check if the secret key is empty
	if len(item) == 0 {
		fmt.Println("Error: Key environment variable not set")
		os.Exit(1)
	}
	return item
}
