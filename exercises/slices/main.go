package main

import "fmt"

func main() {
	sativa, indica := []string{"x", "y"}, []string{"z", "w"}

	fmt.Println("Sativas:")
	for _, v := range sativa {
		fmt.Printf("%v\n", v)
	}
	fmt.Println("Indicas:")
	for _, z := range indica {
		fmt.Printf("%v\n", z)
	}

}
