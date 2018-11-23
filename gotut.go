package main

import (
	"fmt"
)

const usixteenbitmax float64 = 65535
const kmh_multiple float64 = 1.60934 // 1mileh = 1.60934 kmh

type car struct {
	gas_pedal      uint16 // min 0 max 65535
	brake_pedal    uint16
	steering_wheel int16
	top_speed_kmh  float64
}

func (c car) kmh() float64 { // Method: Value Receiver.
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax)
}

func (c car) mph() float64 { // Method: Value Receiver.
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax / kmh_multiple)
}

func (c *car) new_top_speed(newspeed float64) { // Method: Pointer Receiver.
	c.top_speed_kmh = newspeed
}

func car_update(c car, speed float64) car { // Function: --Efficient
	c.top_speed_kmh = speed
	return c
}

func main() {
	a_car := car{22341, 0, 12561, 225.0}
	fmt.Println(a_car.gas_pedal)

	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())

	a_car.new_top_speed(500)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())

	a_car = car_update(a_car, 600)
	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())
}

/*
Methods in go:
1. Value receivers: Calculations given some values. - Makes a copy - Small structures fast operations.
2. Pointer receivers: Modify values. - Modify the original - Big structures heavy operations.
*/
