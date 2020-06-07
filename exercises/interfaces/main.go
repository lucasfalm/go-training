package main

import "fmt"

type (
	vehicleInterface interface {
		turnOn()
	}

	car struct {
		name  string
		color string
	}

	moto struct {
		name  string
		color string
	}
)

func main() {
	var (
		c, m = car{name: "Golf", color: "Black"}, moto{name: "Bis", color: "Blue"}
	)
	var list []vehicleInterface
	list = append(list, c, m)

	for _, y := range list {
		y.turnOn()
	}
}

func turnOn(v vehicleInterface) {
	v.turnOn()
}
func (c car) turnOn() {
	fmt.Println("Car on!")
}

func (m moto) turnOn() {
	fmt.Println("Moto on!")
}
