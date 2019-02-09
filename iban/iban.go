package iban

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

const ibanMaxLength = 34

var letterOffset = int('a') - 10
var validCharacters = regexp.MustCompile("^[0-9A-Z]+$")

func Validate(iban string) (bool, string, error) {
	sanitized := SanitizeInput(iban)

	if len(sanitized) > ibanMaxLength {
		message := fmt.Sprintf("IBAN cannot be longer than 34 characters")
		return false, message, errors.New(message)
	}

	if !CharactersAreValid(sanitized) {
		message := "IBAN contains invalid characters"
		return false, message, errors.New(message)
	}

	countryCode := sanitized[:2]
	country, countryExists := countries[strings.ToUpper(countryCode)]
	if !countryExists {
		message := "Invalid or unsupported country code"
		return false, message, fmt.Errorf("%s (%s)", message, countryCode)
	}

	if len(sanitized) != country.CharCount {
		message := fmt.Sprintf("IBAN length is invalid, expected length is %d, got %d", country.CharCount, len(sanitized))
		return false, message, errors.New(message)
	}

	return true, "IBAN is valid", nil
}

func SanitizeInput(input string) string {
	noWhitespace := strings.Join(strings.Fields(input), "")
	return strings.ToUpper(noWhitespace)
}

func CharactersAreValid(iban string) bool {
	return validCharacters.MatchString(iban)
}

func RuneToIBANInt(char rune) int {
	return int(char) - letterOffset
}
