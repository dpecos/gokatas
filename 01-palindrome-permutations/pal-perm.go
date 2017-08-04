package kata01

import (
	"strings"

	"github.com/dpecos/gokatas/utils/unique"
)

func isPalindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

func permutations(input string) []string {
	var result []string

	for i := 0; i < len(input); i++ {
		current := input[i]
		rest := input[0:i] + input[i+1:]

		partialRes := []string{string(current)}

		if len(rest) == 0 {
			return partialRes
		}

		permRest := permutations(rest)
		for _, perm := range permRest {
			tmp := append(partialRes, perm)
			result = append(result, strings.Join(tmp, ""))
		}
	}

	return unique.Strings(result)
}

func PalindromePermutations(input string) bool {
	for _, perm := range permutations(input) {
		if isPalindrome(perm) {
			return true
		}
	}
	return false
}
