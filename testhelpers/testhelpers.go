package testhelpers

func ErrorsMatch(expected, actual error) bool {
	// Covers both being nil
	if expected == actual {
		return true
	}
	if expected != nil && actual == nil {
		return false
	}
	if expected == nil && actual != nil {
		return false
	}
	return expected.Error() == actual.Error()
}
