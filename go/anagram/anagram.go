package anagram

import (
	"maps"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	result := []string{}

	normalizedSubject := strings.ToLower(subject)
	subjectRunes := []rune(normalizedSubject)
	subjectCount := countRunes(normalizedSubject)

	for _, candidate := range candidates {
		normalizedCandidate := strings.ToLower(candidate)

		if len(subjectRunes) != len([]rune(normalizedCandidate)) {
			continue
		}

		if normalizedSubject == normalizedCandidate {
			continue
		}

		if maps.Equal(subjectCount, countRunes(normalizedCandidate)) {
			result = append(result, candidate)
		}
	}

	return result
}

func countRunes(word string) map[rune]int {
	count := make(map[rune]int)

	for _, r := range word {
		count[r]++
	}

	return count
}
