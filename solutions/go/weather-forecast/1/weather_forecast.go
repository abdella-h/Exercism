// Package weather provides weather forecasting for cities in Goblinocus.
package weather

var (
    // CurrentCondition holds the current weather condition (e.g., "rainy", "sunny").
    // It is used to store the forecast result for a city.
	CurrentCondition string
    
    //CurrentLocation holds the name of the city whose weather is being forecast.
    // Set this before calling Forecast to specify the target city.
	CurrentLocation  string
)
// Forecast returns a string describing the current weather for a given city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
