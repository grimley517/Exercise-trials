// Package weather provides tools to find the forecast for a location.
package weather

var (
    // CurrentCondition represents the summary of the weather conditions.
	CurrentCondition string
    // CurrentLocation represents a location such as a City or Town.
	CurrentLocation  string
)

// Forecast takes a location (city), and a weather condition, and returns a string describing the weather.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
