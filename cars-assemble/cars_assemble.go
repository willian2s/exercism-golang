package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	if successRate < 0 || successRate > 100 {
		panic("successRate must be between 0 and 100")
	}

	successPercentage := successRate / 100
	successfulCars := float64(productionRate) * successPercentage

	return successfulCars
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	carsPerMinute := CalculateWorkingCarsPerHour(productionRate, successRate) / 60
	return int(carsPerMinute)
}

const costPerGroup = 95000
const costPerIndividual = 10000
const carsPerGroup = 10

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	if carsCount < 0 {
		panic("carsCount must be greater than or equal to 0")
	}

	numGroup := carsCount / carsPerGroup
	remainingCars := carsCount % carsPerGroup

	totalCost := uint(numGroup)*costPerGroup + uint(remainingCars)*costPerIndividual

	return totalCost
}
