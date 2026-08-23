package luhn

import "strings"

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")

	if len(id) <= 1 {
		return false
	}

	sum := 0
	double := false

	for i := len(id) - 1; i >= 0; i-- {
		if id[i] < '0' || id[i] > '9' {
			return false
		}

		value := int(id[i] - '0')

		if double {
			value *= 2

			if value > 9 {
				value -= 9
			}
		}

		sum += value
		double = !double
	}

	return sum%10 == 0
}
