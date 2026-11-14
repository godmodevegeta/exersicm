package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var workingCars float64
	workingCars = float64(productionRate) * successRate / 100
	return workingCars
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var workingCarsPerHour float64 = CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(workingCarsPerHour)/60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	var carBunch int = carsCount / 10
	var cars int = carsCount % 10
	return uint(carBunch * 95000 + cars * 10000)
}
