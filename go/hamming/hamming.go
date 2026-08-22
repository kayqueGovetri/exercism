package hamming

import "fmt"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("strings have different lengths: %d and %d", len(a), len(b))
	}
	distance := 0
	for index := range a {
		if a[index] != b[index] {
			distance++
		}
	}
	return distance, nil
}
