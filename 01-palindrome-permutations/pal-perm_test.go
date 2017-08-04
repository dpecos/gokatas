package kata01

import "testing"

func TestPalindromePermutation(t *testing.T) {
	
	var tests = []struct {
		input    string
		expected bool
	}{
		{"a", true},
		{"aa", true},
		{"aba", true},
		{"aab", true},
		{"acb", false},
	}

	for _, test := range tests {
		res := PalindromePermutations(test.input)
		if res != test.expected {
			t.Errorf("PalindromePermutations(%s) => got %t expected %t", test.input, res, test.expected)
		}
	}
}
