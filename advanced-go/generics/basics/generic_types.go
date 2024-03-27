package basics

import "fmt"

func GenericTypes[T any, V comparable](something T, otherSomethingButComparable V) {
	fmt.Printf("this can be anything: %v\n", something)

	fmt.Printf("but this can only be something that is comparable: %v\n", otherSomethingButComparable)
}
