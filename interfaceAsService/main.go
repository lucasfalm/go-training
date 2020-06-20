package main

import (
	"fmt"
)

type dog struct{}

var (
	DogService animalInterface = &dog{} // intancie a variable type animalInterface to export from package
)

type animalInterface interface {
	walk()
}

func main() {
	DogService.walk()
	// DogService.sleep() // doest not work :)
}

// private function
func (d dog) walk() {
	fmt.Println("Uolf uolf")
}

// also private
func (d dog) sleep() {
	fmt.Println("sleeping")
}
