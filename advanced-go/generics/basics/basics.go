package basics

import "fmt"

// NOTE: the "~" makes it accept the type and all the other types that underlay the type
//
//	a.k.a myIntAlias, which is another type not specified here, but with the underling
//	type being one of the listed (int);
type showItParams interface {
	~int32 | ~float64 | ~int
}

type myIntAlias int

func ShowIt[T showItParams](numberOne, numberTwo T) {
	fmt.Printf("adding %v + %v using generics = %v\n", numberOne, numberTwo, addX(numberOne, numberTwo))

	var intAlias myIntAlias = 1

	fmt.Printf("this just works because of the til (~) = %v + %v = %v\n", intAlias, intAlias, addX(intAlias, intAlias))
}

// NOTE: it accepts both int32 or float64
//
//	two ways of defining the type:
//	1. like this
//	2. or using the showItParams which has the type constraint
func addX[T ~int32 | ~float64 | ~int](a, b T) T {
	return a + b
}
