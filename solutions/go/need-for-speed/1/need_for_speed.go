package speed

// TODO: define the 'Car' type struct
type Car struct {
	battery int
	batteryDrain int
	speed int
	distance int
}

type Track struct {
	distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car {
		speed: speed,
		batteryDrain: batteryDrain,
		battery: 100,
		distance: 0,
	}
}

// TODO: define the 'Track' type struct

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	batteryDrain := car.batteryDrain
	can_not_move := car.battery - batteryDrain < 0
	new_distance := car.distance + car.speed
	if !can_not_move {	
		car.battery = car.battery - batteryDrain
		car.distance = new_distance
	}
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	laps := track.distance/car.speed
	battery_drain_on_track := laps * car.batteryDrain
	return !(car.battery - battery_drain_on_track < 0)
}
