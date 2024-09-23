// Package weather provides tool to format forecast message.
package weather

// CurrentCondition represents state of weather.
var CurrentCondition string

// CurrentLocation mean location of forecast.
var CurrentLocation string

// Forecast receive city and condition vars to return formated message for forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
