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
		code   interface{} // empty interface, it could be string, or int, or empty
	}
)

func main() {
	var (
		c, a = car{name: "Golf", color: "Black"}, airplane{number: "Bis", horses: "Blue"}
		list []vehicleInterface
	)
	list = append(list, c, a)
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

func (a airplane) turnOn() {
	fmt.Println("Airplane on!")
}

func stop(v []vehicleInterface) {
	// Use interface to manipulate what type is on the fly
	type carCheck interface {
		vrumvrum()
	}

	// List received
	for _, it := range v {
		// check the type of interface
		switch e := it.(type) {
		case car:
			e.vrumvrum()
		default:
			continue
		}
		// Use interface as a method
		//c, ok := it.(carCheck)
		//if !ok {
		//	continue
		//}

		//c.vrumvrum()
	}

}

func (c car) vrumvrum() {
	fmt.Printf("vrum \n")
}
