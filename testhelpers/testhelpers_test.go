package testhelpers_test

import (
	"errors"
	"testing"

	"github.com/kristofferostlund/iban-validator/testhelpers"
)

func TestErrorsMatch(t *testing.T) {
	sameError := errors.New("I am same")

	cases := []struct {
		a        error
		b        error
		expected bool
	}{
		{nil, nil, true},
		{sameError, sameError, true},
		{errors.New("Oh no"), nil, false},
		{nil, errors.New("Oh no"), false},
		{errors.New("Oh no"), errors.New("Oh no"), true},
	}

	for _, test := range cases {
		actual := testhelpers.ErrorsMatch(test.a, test.b)
		if test.expected != actual {
			t.Errorf("ErrorsMatch(%v, %v), expected: %v, %v", test.a, test.b, test.expected, actual)
		}
	}
}
