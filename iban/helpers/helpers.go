package helpers

import (
	"regexp"
	"strconv"
	"strings"
)

var letterOffset = int('A') - 10
var validCharacters = regexp.MustCompile("^[0-9A-Z]+$")
var numbers = "1234567890"

func SanitizeInput(input string) string {
	noWhitespace := strings.Join(strings.Fields(input), "")
	return strings.ToUpper(noWhitespace)
}

func CharactersAreValid(input string) bool {
	return validCharacters.MatchString(input)
}

func IBANToIntegerString(input string) string {
	integerString := ""
	for _, char := range input {
		if strings.ContainsRune(numbers, char) {
			integerString += string(char)
			continue
		}
		integerString += strconv.Itoa(runeToIBANInt(char))
	}

	return integerString
}

func runeToIBANInt(char rune) int {
	return int(char) - letterOffset
}
