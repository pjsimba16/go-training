// Package weather package initializes current condition and current location variables and runs the forecast function which returns a string containing the initialized variables.
package weather

// CurrentCondition variable of type string that contains the current weather condition in the area.
var CurrentCondition string

// CurrentLocation variable of type string that contains the current location.
var CurrentLocation string

// Forecast function takes in city and condition string inputs and returns a string describing the current weather condition and location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
