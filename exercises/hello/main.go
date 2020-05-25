package main

import (
	"fmt"
)

func main() {
	var (
		age     int
		lastAge int
		name    string
		phrase  = "Ola mundo!"
	)
	fmt.Printf("%q\n", phrase)
	fmt.Printf("%T\n", phrase)

	age, lastAge, name = 22, 21, "Lucas"
	fmt.Println(name, age)

	age, lastAge = lastAge, age

	fmt.Println(age)
}
