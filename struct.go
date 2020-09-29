package main

import (
	"fmt"
)

const uSixteenbitMax float64 = 65535
const kmPerHourMultiple float64 = 1.60934

type car struct {
	gasPedal      int //min:0, max:65535
	brakePedal    int //min:0, max:65535
	steeringWheel int //min:0, max:65535
	topSpeed      float64
}

func (c car) kmh()(float64){
	c.gasPedal = 100
	return float64(c.gasPedal) * (c.topSpeed/uSixteenbitMax)
}

func (c car) mh()(float64){
	return float64(c.gasPedal) * (c.topSpeed/uSixteenbitMax/kmPerHourMultiple)
}

func (c *car) changeTopSpeed(newTopSpeed float64){
	c.topSpeed = newTopSpeed
}

func main() {
	newCar := car{
		gasPedal:      65000,
		brakePedal:    0,
		steeringWheel: 12562,
		topSpeed:      225.0}
	fmt.Println("gas_pedal:", newCar.gasPedal, "brake_pedal:", newCar.brakePedal, "steering_wheel:", newCar.steeringWheel)
	fmt.Println(newCar.kmh())
	fmt.Println(newCar.mh())
	newCar.changeTopSpeed(500)
	fmt.Println(newCar.kmh())
	fmt.Println(newCar.mh())
}
