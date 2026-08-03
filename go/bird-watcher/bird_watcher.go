package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0
	for _, value := range birdsPerDay {
		total += value
	}
	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	total := 0
	start := (week - 1) * 7
	end := week * 7
	for index, value := range birdsPerDay {
		if index >= start && index < end {
			total += value
		}
	}
	return total
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for index, _ := range birdsPerDay {
		if index%2 == 0 {
			birdsPerDay[index] += 1
		}
	}
	return birdsPerDay
}
