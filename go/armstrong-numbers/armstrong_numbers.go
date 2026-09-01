package armstrongnumbers

import (
	"math"
	"strconv"
)

func IsNumber(n int) bool {
	size := strconv.Itoa(n)
	digits := make([]int, len(size))
	result := 0
	for i := range digits {
		value, _ := strconv.Atoi(string(size[i]))
		result += int(math.Pow(float64(value), float64(len(size))))
	}
	return result == n
}
