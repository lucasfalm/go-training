package main

import (
	"fmt"
)

func main() {
	var (
		age  int
		name string
	)
	phrase := "Ola mundo!"
	fmt.Println(phrase)

	age, name = 22, "Lucas"
	fmt.Println(name, age)
}
