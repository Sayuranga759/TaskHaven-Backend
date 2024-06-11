package validator

import (
	"net/mail"
	"regexp"
)

// Validate the filed value is valid for expected regexp
func containsOnly(field, regExpStr string) bool {
	reg, _ := regexp.Compile(regExpStr)
	return !reg.MatchString(field)
}

// Validate the email is valid as expected
func validateEmail(field string) bool {
	_, err := mail.ParseAddress(field)
	return err == nil
}

// Validate the password is valid as expected
func validatePassword(field string, regexes []string) bool {
	for _,regex := range regexes {
		reg, _ := regexp.Compile(regex)
		if !reg.MatchString(field) {
			return false
		}
	}

	return true
}










