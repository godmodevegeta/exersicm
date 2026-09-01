package jedlik

import "fmt"

// TODO: define the 'Drive()' method
func (car *Car) Drive() {
	current_distance := car.distance
	current_battery := car.battery
	new_distance := current_distance + car.speed
	new_battery := current_battery - car.batteryDrain
	if new_battery < 0 {
		return
	}
	car.battery = new_battery
	car.distance = new_distance
}

// TODO: define the 'DisplayDistance() string' method
func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

// TODO: define the 'DisplayBattery() string' method
func (car Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)	
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (car Car) CanFinish(trackDistance int) bool {
	laps := trackDistance/car.speed
	new_battery := car.battery - laps * car.batteryDrain
	if laps * car.speed < trackDistance {
		new_battery = new_battery - car.batteryDrain
	}
	return !(new_battery < 0)
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
