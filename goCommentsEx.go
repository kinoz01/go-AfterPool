// Package weather provides tools to forecast
// weather.
package weather
// CurrentCondition is a string varaible that tell us current weather condition.
var CurrentCondition string
// CurrentLocation is a string varaible that tell us current location.
var CurrentLocation string

// Forecast returns a string value that forecast the current weather condition in a givven city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
