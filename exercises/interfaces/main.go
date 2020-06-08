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

	airplane struct {
		number string
		horses string
	}
)

func main() {
	var (
		c, m = car{name: "Golf", color: "Black"}, airplane{number: "Bis", horses: "Blue"}
		list []vehicleInterface
	)
	list = append(list, c, m)
	stop(list)

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

func (m airplane) turnOn() {
	fmt.Println("Airplane on!")
}

func stop(v []vehicleInterface) {
	// Use interface to manipulate what type is on the fly
	type carCheck interface {
		vrumvrum()
	}
	// List received
	for _, it := range v {
		c, ok := it.(carCheck)
		if !ok {
			continue
		}

		c.vrumvrum()
	}

}

func (c car) vrumvrum() {
	fmt.Printf("vrum \n")
}
