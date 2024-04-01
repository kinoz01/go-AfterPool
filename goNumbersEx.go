package main

import "fmt"

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate)*successRate/100
	// panic("CalculateWorkingCarsPerHour not implemented")
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	x :=  CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(x/60)
	// panic("CalculateWorkingCarsPerMinute not implemented")
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	return uint((carsCount%10)*10000+(carsCount/10)*95000)
	// panic("CalculateCost not implemented")
}

func main() {
	x := CalculateWorkingCarsPerHour(1547, 25.5)
	fmt.Println("Working cars per hour with 1547 prodution rate and 90 sucess rate:", x)
	fmt.Println("Working cars per minute with 1547 prodution rate and 90 sucess rate:", int(x/60))
	fmt.Println("cost for 37 cars:", CalculateCost(21))
}
