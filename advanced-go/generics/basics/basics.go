package basics

import "fmt"

type ShowItParameter interface {
	int32 | float64
}

func ShowIt[T ShowItParameter](numberOne, numberTwo T) {
	fmt.Printf("adding %v + %v using generics = %v\n", numberOne, numberTwo, addX(numberOne, numberTwo))
}

// NOTE: it accepts both int32 or float64
//
//	two ways of defining the type:
//	1. like this
//	2. or using the ShowItParameter which has the type constraint
func addX[T int32 | float64](a, b T) T {
	return a + b
}
