package main

import (
	"fmt"
)

func main() {
	var (
		age int
		// lastAge int
		name   string
		phrase = "Ola mundo!"
	)
	fmt.Printf("%q\n", phrase)
	fmt.Printf("%T\n", phrase)

	age, name = 22, "Lucas"

	fmt.Printf("Seu nome é %s, e sua idade é %d\n", name, age)

	fmt.Printf("Sua idade é %[2]v, e seu nome é %[1]v\n", name, age)
	/*
		age, lastAge = lastAge, age

		fmt.Println(age)
	*/
}
