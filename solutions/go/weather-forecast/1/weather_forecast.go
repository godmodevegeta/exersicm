// Package weather provides forecasting tools. 
package weather

var (
// CurrentCondition is input for current condition.
	CurrentCondition string
// CurrentLocation is input for current condition.
	CurrentLocation  string
)

// Forecast method which forecasts.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
