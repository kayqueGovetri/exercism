package bottlesong

import (
	"fmt"
	"strings"
)

var number = map[int]string{
	10: "Ten",
	9:  "Nine",
	8:  "Eight",
	7:  "Seven",
	6:  "Six",
	5:  "Five",
	4:  "Four",
	3:  "Three",
	2:  "Two",
	1:  "One",
	0:  "no",
}

func verse(bottles int) []string {
	bottleWord := "bottles"

	if bottles == 1 {
		bottleWord = "bottle"
	}

	nextBottleWord := "bottles"

	if bottles-1 == 1 {
		nextBottleWord = "bottle"
	}

	return []string{
		fmt.Sprintf(
			"%s green %s hanging on the wall,",
			number[bottles],
			bottleWord,
		),
		fmt.Sprintf(
			"%s green %s hanging on the wall,",
			number[bottles],
			bottleWord,
		),
		"And if one green bottle should accidentally fall,",
		fmt.Sprintf(
			"There'll be %s green %s hanging on the wall.",
			strings.ToLower(number[bottles-1]),
			nextBottleWord,
		),
	}
}

func Recite(startBottles, takeDown int) []string {
	result := []string{}

	for i := 0; i < takeDown; i++ {
		bottles := startBottles - i

		result = append(result, verse(bottles)...)

		if i < takeDown-1 {
			result = append(result, "")
		}
	}

	return result
}
