package main

import (
	"fmt"
)

func main() {
	var (
		age int
		// lastAge int
		name      string
		phrase    = "Ola mundo!"
		daysEarth float64
	)
	fmt.Printf("%q\n", phrase)
	fmt.Printf("%T\n", phrase)

	age, name, daysEarth = 22, "Lucas", 364.56945

	fmt.Printf("Seu nome é %s, e sua idade é %d\n", name, age)

	fmt.Printf("Sua idade é %[2]v, e seu nome é %[1]v\n", name, age)

	fmt.Printf("A quantidade de dias terrestres é %0.f\n", daysEarth)
	fmt.Printf("A quantidade de dias terrestres é %0.1f\n", daysEarth)
	fmt.Printf("A quantidade de dias terrestres é %0.5f\n", daysEarth)
	/*
		age, lastAge = lastAge, age

		fmt.Println(age)
	*/
}
