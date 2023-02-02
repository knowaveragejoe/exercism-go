package elon

import "fmt"

func (car *Car) Drive() {
	if car.battery < car.batteryDrain {
		return
	}

	car.distance += car.speed
	car.battery -= car.batteryDrain
}

func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %v meters", car.distance)
}

func (car Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %v", car.battery) + "%"
}

func (car Car) CanFinish(trackDistance int) bool {
	drives := car.battery / car.batteryDrain
	carRange := car.speed * drives

	if car.battery <= 0 {
		return false
	}

	return carRange >= trackDistance
}
