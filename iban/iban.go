package iban

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/kristofferostlund/iban-validator/iban/helpers"
)

const (
	ibanMaxLength        = 34
	modValue             = 97
	checkDigitsChunkSize = 7
)

func Validate(iban string) (bool, string, error) {
	sanitized := helpers.SanitizeInput(iban)

	if len(sanitized) > ibanMaxLength {
		return false, fmt.Sprintf("IBAN cannot be longer than %d characters", ibanMaxLength), nil
	}

	if !helpers.CharactersAreValid(sanitized) {
		return false, "IBAN contains invalid characters", nil
	}

	countryCode := sanitized[:2]
	country, countryExists := countries[strings.ToUpper(countryCode)]
	if !countryExists {
		return false, "Invalid or unsupported country code", nil
	}

	if len(sanitized) != country.CharCount {
		message := fmt.Sprintf(
			"IBAN length is invalid, expected length is %d, got %d",
			country.CharCount,
			len(sanitized),
		)
		return false, message, nil
	}

	isValid, err := ValidateCheckDigits(sanitized)
	if err != nil {
		message := "Invalid IBAN, could not verify check digits"
		return false, message, fmt.Errorf("%s: %v", message, err)
	}

	if !isValid {
		message := "Invalid IBAN, check digits are invalid for the provided IBAN"
		return false, message, nil
	}

	return true, "IBAN is valid", nil
}

func ValidateCheckDigits(input string) (bool, error) {
	// From https://en.wikipedia.org/wiki/International_Bank_Account_Number#Modulo_operation_on_IBAN
	reordered := input[4:] + input[:4]
	integerString := helpers.IBANToIntegerString(reordered)

	nMod := 0
	current := integerString[:2]
	remainder := integerString[2:]

	for len(remainder) > 0 {
		i := int(math.Min(float64(checkDigitsChunkSize), float64(len(remainder))))

		if nMod == 0 {
			current = current + remainder[:i]
		} else {
			current = strconv.Itoa(nMod) + remainder[:i]
		}
		remainder = remainder[i:]

		N, err := strconv.Atoi(current)
		if err != nil {
			return false, err
		}

		nMod = int(math.Mod(float64(N), float64(modValue)))
	}

	if nMod != 1 {
		return false, nil
	}

	return true, nil
}
