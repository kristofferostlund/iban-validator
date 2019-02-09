package iban_test

import (
	"fmt"
	"testing"

	"github.com/kristofferostlund/iban-validator/iban"
)

func TestValidate_valid(t *testing.T) {
	// Taken from https://www.iban.com/structure
	cases := []struct {
		input string
	}{
		{"se72 8000 0810 3400 0978 3242"},
		{"SE7280000810340009783242"},
		{"GB82 WEST 1234 5698 7654 32	"},
	}

	for _, test := range cases {
		isValid, message, err := iban.Validate(test.input)

		actuals := fmt.Sprintf(
			"got isValid: %t, message: \"%s\", error: \"%v\"",
			isValid,
			message,
			err,
		)

		if !isValid {
			t.Errorf(
				"Validate(%+v), expected IBAN to be valid, %s",
				test.input,
				actuals,
			)
		}

		if err != nil {
			t.Errorf(
				"Validate(%+v), expected IBAN to not return error, %s",
				test.input,
				actuals,
			)
		}
	}
}
func TestValidate_valid_supportedCountries(t *testing.T) {
	// Taken from https://www.iban.com/structure
	cases := []struct {
		country string
		input   string
	}{
		{"Andorra", "AD14 0008 0001 0012 3456 7890"},
		{"Austria", "AT48 3200 0000 1234 5864"},
		{"Azerbaijan", "AZ96 AZEJ 0000 0000 0012 3456 7890"},
		{"Bahrain", "BH02 CITI 0000 1077 1816 11"},
		{"Belarus", "BY86 AKBB 1010 0000 0029 6600 0000"},
		{"Belgium", "BE71 0961 2345 6769"},
		{"Bosnia and Herzegovina", "BA39 3385 8048 0021 1234"},
		{"Brazil", "BR15 0000 0000 0000 1093 2840 814P 2"},
		{"Bulgaria", "BG18 RZBB 9155 0123 4567 89"},
		{"Costa Rica", "CR23 0151 0841 0026 0123 45"},
		{"Croatia", "HR17 2360 0001 1012 3456 5"},
		{"Cyprus", "CY21 0020 0195 0000 3570 0123 4567"},
		{"Czech Republic", "CZ55 0800 0000 0012 3456 7899"},
		{"Denmark", "DK95 2000 0123 4567 89"},
		{"Dominican Republic", "DO22 ACAU 0000 0000 0001 2345 6789"},
		{"Estonia", "EE47 1000 0010 2014 5685"},
		{"Faroe Islands", "FO92 6460 0123 4567 89"},
		{"Finland", "FI14 1009 3000 1234 58"},
		{"France", "FR76 3000 6000 0112 3456 7890 189"},
		{"Georgia", "GE60 NB00 0000 0123 4567 89"},
		{"Germany", "DE75 5121 0800 1245 1261 99"},
		{"Gibraltar", "GI04 BARC 0000 0123 4567 890"},
		{"Greece", "GR96 0810 0010 0000 0123 4567 890"},
		{"Greenland", "GL89 6471 0123 4567 89"},
		{"Guatemala", "GT20 AGRO 0000 0000 0012 3456 7890"},
		{"Hungary", "HU93 1160 0006 0000 0000 1234 5676"},
		{"Iceland", "IS75 0001 1212 3456 3108 9620 99"},
		{"Ireland", "IE64 IRCE 9205 0112 3456 78"},
		{"Israel", "IL17 0108 0000 0001 2612 345"},
		{"Italy", "IT60 X054 2811 1010 0000 0123 456"},
		{"Jordan", "JO71 CBJO 0000 0000 0000 1234 5678 90"},
		{"Kazakhstan", "KZ56 3190 0000 1234 4567"},
		{"Kosovo", "XK05 1212 0123 4567 8906"},
		{"Kuwait", "KW81 CBKU 0000 0000 0000 1234 5601 01"},
		{"Latvia", "LV97 HABA 0012 3456 7891 0"},
		{"Lebanon", "LB92 0007 0000 0000 1231 2345 6123"},
		{"Liechtenstein", "LI74 0880 6123 4567 8901 2"},
		{"Lithuania", "LT60 1010 0123 4567 8901"},
		{"Luxembourg", "LU12 0010 0012 3456 7891"},
		{"Macedonia", "MK07 2000 0278 5123 453"},
		{"Malta", "MT31 MALT 0110 0000 0000 0000 0000 123"},
		{"Mauritania", "MR13 0002 0001 0100 0012 3456 753"},
		{"Mauritius", "MU43 BOMM 0101 1234 5678 9101 000M UR"},
		{"Moldova", "MD21 EX00 0000 0000 0123 4567"},
		{"Monaco", "MC58 1009 6180 7901 2345 6789 085"},
		{"Montenegro", "ME25 5050 0001 2345 6789 51"},
		{"Netherlands", "NL02 ABNA 0123 4567 89"},
		{"Norway", "NO83 3000 1234 567"},
		{"Pakistan", "PK36 SCBL 0000 0011 2345 6702"},
		{"Palestine", "PS92 PALS 0000 0000 0400 1234 5670 2"},
		{"Poland", "PL10 1050 0099 7603 1234 5678 9123"},
		{"Portugal", "PT50 0027 0000 0001 2345 6783 3"},
		{"Qatar", "QA54 QNBA 0000 0000 0000 6931 2345 6"},
		{"Romania", "RO09 BCYP 0000 0012 3456 7890"},
		{"San Marino", "SM76 P085 4009 8121 2345 6789 123"},
		{"Saudi Arabia", "SA44 2000 0001 2345 6789 1234"},
		{"Serbia", "RS35 1050 0812 3123 1231 73"},
		{"Slovak Republic", "SK89 7500 0000 0000 1234 5671"},
		{"Slovenia", "SI56 1920 0123 4567 892"},
		{"Spain", "ES79 2100 0813 6101 2345 6789"},
		{"Sweden", "SE72 8000 0810 3400 0978 3242"},
		{"Switzerland", "CH56 0483 5012 3456 7800 9"},
		{"Timor-Leste", "TL38 0010 0123 4567 8910 106"},
		{"Tunisia", "TN59 0401 8104 0049 4271 2345"},
		{"Turkey", "TR32 0010 0099 9990 1234 5678 90"},
		{"Ukraine", "UA90 3052 9929 9000 4149 1234 5678 9"},
		{"United Arab Emirates", "AE46 0090 0000 0012 3456 789"},
		{"United Kingdom", "GB33 BUKB 2020 1555 5555 55"},
		{"Virgin Islands, British", "VG21 PACG 0000 0001 2345 6789"},
	}

	for _, test := range cases {
		isValid, message, err := iban.Validate(test.input)

		actuals := fmt.Sprintf(
			"got isValid: %t, message: \"%s\", error: \"%v\"",
			isValid,
			message,
			err,
		)

		if !isValid {
			t.Errorf(
				"Validate(%+v), expected IBAN (%s) to be valid, %s",
				test.input,
				test.country,
				actuals,
			)
		}

		if err != nil {
			t.Errorf(
				"Validate(%+v), expected IBAN (%s) to not return error, %s",
				test.input,
				test.country,
				actuals,
			)
		}
	}
}

func TestValidate_unsupportedCountries(t *testing.T) {
	// Taken from https://www.iban.com/structure
	cases := []struct {
		country     string
		countryCode string
		input       string
	}{
		{"El Salvador", "SV", "SV43 ACAT 0000 0000 0000 0012 3123"},
		{"Iraq", "IQ", "IQ20 CBIQ 8618 0010 1010 500"},
		{"Saint Lucia", "LC", "LC14 BOSL 1234 5678 9012 3456 7890 1234"},
		{"Sao Tome and Principe", "ST", "ST23 0002 0000 0289 3557 1014 8"},
		{"Seychelles", "SC", "SC52 BAHL 0103 1234 5678 9012 3456 USD"},
	}

	for _, test := range cases {
		isValid, message, err := iban.Validate(test.input)

		actuals := fmt.Sprintf(
			"got isValid: %t, message: \"%s\", error: \"%v\"",
			isValid,
			message,
			err,
		)

		if isValid {
			t.Errorf(
				"Validate(%+v), expected IBAN (from %s) to be invalid, %s",
				test.input,
				test.country,
				actuals,
			)
		}

		if err != nil {
			t.Errorf(
				"Validate(%+v), expected IBAN (from %s) not return an error, %s",
				test.input,
				test.country,
				actuals,
			)
		}

		expectedMessage := "Invalid or unsupported country code"
		if expectedMessage != message {
			t.Errorf(
				"Validate(%+v), expected IBAN (from %s) to return the message \"%s\", %s",
				test.input,
				test.country,
				expectedMessage,
				actuals,
			)
		}
	}
}

func TestValidate_invalid_noError(t *testing.T) {
	// Taken from https://www.iban.com/structure
	cases := []struct {
		input   string
		message string
	}{
		{"SE72 8000 0810 3400 0978 3242 8000 0810 3400 0978 3242 8000 0810 3400 0978 3242", "IBAN cannot be longer than 34 characters"},
		{"🏦🏦🏦🏦", "IBAN contains invalid characters"},
		{"SE72 8000 0810 3400 0978 f̅f̅f̅f̅", "IBAN contains invalid characters"},
		{"SV43 ACAT 0000 0000 0000 0012 3123", "Invalid or unsupported country code"},
		{"SE72 8000 0810 3400 0978 3242 3242", "IBAN length is invalid, expected length is 24, got 28"},
		// {"SE72 8000 0810 3400 0978 3241", "Invalid IBAN, could not verify check digits"}, // This error should not be able to occur
		{"SE72 8000 0810 3400 0978 3241", "Invalid IBAN, check digits are invalid for the provided IBAN"},
	}

	for _, test := range cases {
		isValid, message, err := iban.Validate(test.input)

		actuals := fmt.Sprintf(
			"got isValid: %t, message: \"%s\", error: \"%v\"",
			isValid,
			message,
			err,
		)

		if isValid {
			t.Errorf(
				"Validate(%+v), expected IBAN to be invalid, %s",
				test.input,
				actuals,
			)
		}

		if err != nil {
			t.Errorf(
				"Validate(%+v), expected IBAN to return an error, %s",
				test.input,
				actuals,
			)
		}

		if test.message != message {
			t.Errorf(
				"Validate(%+v), expected IBAN to return the message \"%+s\", %s",
				test.input,
				test.message,
				actuals,
			)
		}
	}
}

func TestRuneToIBANInt(t *testing.T) {
	cases := []struct {
		input    rune
		expected int
	}{
		{'A', 10},
		{'B', 11},
		{'Z', 35},
	}

	for _, test := range cases {
		actual := iban.RuneToIBANInt(test.input)
		if test.expected != actual {
			t.Errorf("RuneToIBANInt(%+v), expected %d, got: %d", test.input, test.expected, actual)
		}
	}
}

func TestSanitizeInput(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"1234 5678 90AB CDEF", "1234567890ABCDEF"},
		{"1234 5678 90ab cdef", "1234567890ABCDEF"},
	}

	for _, test := range cases {
		actual := iban.SanitizeInput(test.input)
		if test.expected != actual {
			t.Errorf("SanitizeInput(%+v), expected %s, got: %s", test.input, test.expected, actual)
		}
	}
}

func TestCharactersAreValid(t *testing.T) {
	cases := []struct {
		input    string
		expected bool
	}{
		{"1234567890ABCDEF", true},
		{"HELLOTHERE", true},
		{"1234 5678 90AB CDEF", false},
		{"", false},
	}

	for _, test := range cases {
		actual := iban.CharactersAreValid(test.input)
		if test.expected != actual {
			t.Errorf("CharactersAreValid(%+v), expected %t, got: %t", test.input, test.expected, actual)
		}
	}
}

func TestValidateCheckDigits_valid(t *testing.T) {
	input := "GB82WEST12345698765432"

	isValid, err := iban.ValidateCheckDigits(input)

	if !isValid {
		t.Errorf("CalculateChecksumString(%+v), expected IBAN to be valid, got isValid: %t, error: %v", input, isValid, err)
	}
}
