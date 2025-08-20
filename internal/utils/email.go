package utils

import (
	"regexp"
)

func CheckEmail(email string) bool {
	// Check if email is empty
	if email == "" {
		return false
	}

	// Validate email format
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return false
	}
	return true
}
