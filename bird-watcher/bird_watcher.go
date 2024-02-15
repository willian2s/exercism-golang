package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var birds int

	for i := 0; i < len(birdsPerDay); i++ {
		birds += birdsPerDay[i]
	}
	return birds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	daysPerWeek := 7

	startIndex := (week - 1) * daysPerWeek
	endIndex := startIndex + daysPerWeek

	weekTotal := 0
	for i := startIndex; i < endIndex; i++ {
		weekTotal += birdsPerDay[i]
	}
	return weekTotal
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i] += 1
	}
	return birdsPerDay
}
