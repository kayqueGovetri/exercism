package isogram

import "strings"

func IsIsogram(word string) bool {
	existsLetter := make(map[rune]bool)
	for _, letter := range strings.ToLower(word) {
		if letter == '-' || letter == ' ' {
			continue
		}
		if existsLetter[letter] {
			return false
		}
		existsLetter[letter] = true
	}
	return true
}
