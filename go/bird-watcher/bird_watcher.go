package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var totalCount int
	for _, v := range birdsPerDay {
		totalCount += v
	}

	return totalCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var offset int = 7 * (week - 1)
	return TotalBirdCount(birdsPerDay[offset : offset+7])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i, v := range birdsPerDay {
		if i%2 == 0 {
			birdsPerDay[i] = v + 1
		}
	}

	return birdsPerDay
}
