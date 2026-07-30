// Package weather provides tools to get forecast.
package weather

var (
	//CurrentCondition represent actual condition.
	CurrentCondition string
	//CurrentLocation represent a city.
	CurrentLocation string
)

// Forecast returns an string value to the information forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
