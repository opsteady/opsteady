package utils

// UniqueNonEmptyElementsOf makes the string list unique
func UniqueNonEmptyElementsOf(s []string) []string {
	if len(s) == 0 {
		return []string{}
	}

	unique := make(map[string]bool)
	for _, elem := range s {
		if elem != "" {
			unique[elem] = true
		}
	}

	var uniqueStrings = []string{}
	for elem := range unique {
		uniqueStrings = append(uniqueStrings, elem)
	}

	return uniqueStrings
}

// ReverseStringArray reverses the array
func ReverseStringArray(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}
