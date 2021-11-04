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
