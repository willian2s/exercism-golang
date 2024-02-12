// Package weather return the current forecast.
package weather

// CurrentCondition is a current condition.
var CurrentCondition string

// CurrentLocation is a current location to forecast.
var CurrentLocation string

// Forecast returns the current weather condition for the given city.
//
// Parameters:
//
//	city string - the name of the city for which the weather condition is requested.
//	condition string - the current weather condition of the city.
//
// Return type:
//
//	string - the current weather condition of the city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
