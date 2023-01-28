// Package weather - displays weather info for Goblinocus.
package weather

// CurrentCondition current conditions.
var CurrentCondition string

// CurrentLocation current location.
var CurrentLocation string

// Forecast - displays the forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
