package main

import (
	"fmt"
)

func main() {
	type myType int32
	var b myType

	b = 5
	// use constant untyped, thats why I dont need to convert in myType type
	const a = 1

	result := a + b

	fmt.Printf("Resultado %v\n", result)

	c := 2
	// use conversion type
	result *= myType(c)
	fmt.Printf("Resultado %v", result)

}
