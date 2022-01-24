package elon

import "strconv"

// TODO: define the 'Drive()' method
func (car *Car) Drive() {
	if car.battery >= car.batteryDrain {
		car.distance += car.speed
		car.battery -= car.batteryDrain
	}
}

// TODO: define the 'DisplayDistance() string' method
func (car *Car) DisplayDistance() string {
	return "Driven " + strconv.Itoa(car.distance) + " meters"
}

// TODO: define the 'DisplayBattery() string' method
func (car *Car) DisplayBattery() string {
	return "Battery at " + strconv.Itoa(car.battery) + "%"
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (car *Car) CanFinish(trackDistance int) bool {
	return trackDistance <= (car.battery / car.batteryDrain * car.speed)
}
