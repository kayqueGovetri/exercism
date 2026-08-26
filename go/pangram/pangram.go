package pangram

import "strings"

func IsPangram(input string) bool {
	letters := make(map[rune]bool)
	for _, char := range strings.ToLower(input) {
		if char >= 'a' && char <= 'z' {
			letters[char] = true
		}
	}
	return len(letters) == 26
}
