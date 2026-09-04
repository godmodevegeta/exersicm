package meteorology

import (
	"fmt"
)

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (temperatureUnit TemperatureUnit) String() string {
	if temperatureUnit == 0 {
		return fmt.Sprintf("°C")
	}	
	return fmt.Sprintf("°F")
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (temperature Temperature) String() string {
	return fmt.Sprintf("%d %s", temperature.degree, temperature.unit)
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (speedunit SpeedUnit) String() string {
	unit := map[SpeedUnit]string{
		0: "km/h",
		1: "mph",
	}
	return unit[speedunit]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed

func (speed Speed) String() string {
	return fmt.Sprintf("%d %s", speed.magnitude, speed.unit.String())
}


type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData

func (meteorology MeteorologyData) String() string {
	return fmt.Sprintf("%s: %s, Wind %s at %s, %d%% Humidity", meteorology.location, meteorology.temperature, meteorology.windDirection, meteorology.windSpeed, meteorology.humidity)
}
