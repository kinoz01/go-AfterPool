package main

import "fmt"

type Car struct {
	battery, batteryDrain, speed, distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:      100,
		speed:        speed, 
		batteryDrain: batteryDrain,
	}
}

type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{distance: distance,}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if car.battery >= car.batteryDrain  {
		car.battery -= car.batteryDrain
		car.distance += car.speed
		return car
	}
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	for i := car.battery; i > 0; i -= car.batteryDrain {
		car = Drive(car)
		Drive(car)
	}
	return car.distance >= track.distance 

	// return (car.battery*car.speed/car.batteryDrain) >= track.distance
	/*
	car.batteryDrain ------> car.speed
	car.battery -----------> maxDistance
	*/
}

func main() {
	car := NewCar(5, 2)
	fmt.Println(car)

	distance := 800
	track := NewTrack(distance)
	fmt.Println(track)

	car	= NewCar(8, 3)
	car = Drive(car)
	fmt.Println(car)

	speed := 5
	batteryDrain := 2
	car = NewCar(speed, batteryDrain)
	distance = 100
	track = NewTrack(distance)
	fmt.Println(CanFinish(car, track))
}

// Need for Speed