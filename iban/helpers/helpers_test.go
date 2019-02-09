package helpers_test

import (
	"testing"

	"github.com/kristofferostlund/iban-validator/iban/helpers"
)

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
		actual := helpers.RuneToIBANInt(test.input)
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
		actual := helpers.SanitizeInput(test.input)
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
		actual := helpers.CharactersAreValid(test.input)
		if test.expected != actual {
			t.Errorf("CharactersAreValid(%+v), expected %t, got: %t", test.input, test.expected, actual)
		}
	}
}
