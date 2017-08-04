package kata01

import "testing"

func TestIsPalindrome(t *testing.T) {

	var tests = []struct {
		input    string
		expected bool
	}{
		{"a", true},
		{"aa", true},
		{"aba", true},
		{"aab", false},
	}

	for _, test := range tests {
		res := isPalindrome(test.input)
		if res != test.expected {
			t.Errorf("isPalindrome(%s) => was %t but expected %t", test.input, res, test.expected)
		}
	}
}
