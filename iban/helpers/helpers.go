package helpers

import (
	"regexp"
	"strings"
)

var letterOffset = int('A') - 10
var validCharacters = regexp.MustCompile("^[0-9A-Z]+$")

func SanitizeInput(input string) string {
	noWhitespace := strings.Join(strings.Fields(input), "")
	return strings.ToUpper(noWhitespace)
}

func CharactersAreValid(input string) bool {
	return validCharacters.MatchString(input)
}

func RuneToIBANInt(char rune) int {
	return int(char) - letterOffset
}
