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

func (c car) kmh() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax)
}

func (c car) mph() float64 {
	return float64(c.gas_pedal) * (c.top_speed_kmh / usixteenbitmax / kmh_multiple)
}

func main() {
	a_car := car{22341, 0, 12561, 225.0}
	fmt.Println(a_car.gas_pedal)

	fmt.Println(a_car.kmh())
	fmt.Println(a_car.mph())
}

/*
Methods in go:
1. Value recivers: Calculations given some values.
2. Pointer recivers: Modify values.
*/
