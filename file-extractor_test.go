package main

import "testing"

func TestRremoveWhiteSpace(t *testing.T) {
	tests := map[string]string{
		"   tests": "tests",
		"tests":    "tests",
		" \ttests": "tests",
	}

	for input, result := range tests {
		if removeWhiteSpace([]byte(input)) != result {
			t.Errorf("Mismatch %s should be %s", input, result)
		}
	}
}
