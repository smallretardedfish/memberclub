package server

import (
	"net/mail"
	"regexp"
)

type ValidationPair struct {
	CredToBeValidated string
	Value             string
}

func Validation(pair ValidationPair) bool {
	name := regexp.MustCompile(`^([a-zA-Z]{2,}\s[a-zA-Z]{1,}'?-?[a-zA-Z]{2,}\s?([a-zA-Z]{1,})?)`)

	switch pair.CredToBeValidated {
	case "email":
		_, err := mail.ParseAddress(pair.Value)
		if err != nil {
			return false
		}
	case "name":
		return name.MatchString(pair.Value)
	default:
		return false
	}
	return true
}
