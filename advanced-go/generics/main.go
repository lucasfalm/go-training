package main

import "github.com/lucasfalm/go-training/advanced-go/generics/basics"

func main() {
	basics.ShowIt(int32(2), int32(938))

	// NOTE: I don't have to cast here because I'm using the package constraints from go
	// 			 which defines all integer types possible
	basics.PackageConstraints(5, 92)

	// NOTE: similar here, but instead my function accepts any type as the first parameter
	// 			 and the second can be anything we can use comparison (>, <, ==)
	basics.GenericTypes("anything!", 89)
}
