package kata01

import "testing"
import "reflect"

func TestPermuations(t *testing.T) {
	var tests = []struct {
		input    string
		expected []string
	}{
		{"a", []string{"a"}},
		{"ab", []string{"ab", "ba"}},
		{"aab", []string{"aab", "aba", "baa"}},
	}

	for _, test := range tests {
		res := permutations(test.input)
		if !reflect.DeepEqual(res, test.expected) {
			t.Errorf("permutations(%s) => was %v but expected %v", test.input, res, test.expected)
		}
	}

}
