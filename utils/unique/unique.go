package unique

func Strings(input []string) []string {
	result := make([]string, 0, len(input))
	m := make(map[string]bool)

	for _, str := range input {
		if _, ok := m[str]; !ok {
			m[str] = true
			result = append(result, str)
		}
	}

	return result
}
