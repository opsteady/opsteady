package utils

// UniqueNonEmptyElementsOf makes the string list unique
func UniqueNonEmptyElementsOf(s []string) []string {
	unique := make(map[string]bool, len(s))
	uniqueStrings := make([]string, 0)

	for _, elem := range s {
		if len(elem) != 0 {
			if !unique[elem] {
				uniqueStrings = append(uniqueStrings, elem)
				unique[elem] = true
			}
		}
	}

	return uniqueStrings
}

// ReverseStringArray reverses the array
func ReverseStringArray(s []string) []string {
	newArray := make([]string, len(s))
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		newArray[i], newArray[j] = s[j], s[i]
	}
	return newArray
}
