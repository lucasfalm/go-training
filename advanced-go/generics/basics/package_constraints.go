package basics

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// NOTE: there is the constraints package from go that already has interfaces defining
//
//	type constraints (as Integer, Float, Complex, etc)
func PackageConstraints[T constraints.Integer | constraints.Float](numberOne, numberTwo T) {
	fmt.Printf(("using package constraints: %v - %v = %v"), numberOne, numberTwo, subX(numberOne, numberTwo))
}

func subX[T constraints.Integer | constraints.Float](numberOne, numberTwo T) T {
	return numberOne - numberTwo
}
