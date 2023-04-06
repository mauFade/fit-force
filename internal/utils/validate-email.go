package utils

import (
	"regexp"
)

func ValidateMail(email string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	return regex.MatchString(email)
}
