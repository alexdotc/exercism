// Package weather forecasts the current weather condition of various cities in Goblinocus.
package weather

// CurrentCondition is the weather condition for the last forecast.
var CurrentCondition string
// CurrentLocation is the city for the last forecast.
var CurrentLocation string

// Forecast updates the weather forecast for the given city to the given condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
