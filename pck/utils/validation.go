package utils

import (
	"fmt"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

func ValidatePassword(password string) bool {
	// Length check
	if len(password) < 8 {
		return false
	}

	// Regex patterns for different character types
	upperCase := regexp.MustCompile(`[A-Z]`)
	lowerCase := regexp.MustCompile(`[a-z]`)
	digit := regexp.MustCompile(`\d`)
	specialChar := regexp.MustCompile(`[@$!%*?&]`)

	// Validate all conditions
	return upperCase.MatchString(password) &&
		lowerCase.MatchString(password) &&
		digit.MatchString(password) &&
		specialChar.MatchString(password)
}

func HashPassword(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error in hashing the password : %w", err)
	}
	return string(hashedPass), nil

}

func ComparePaswords(password, hasspassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hasspassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func ValidateEmail(email string) bool {
	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
