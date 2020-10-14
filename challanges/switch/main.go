package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		name = os.Args[1]
	)

	switch name {
	case "Lucas":
		fmt.Println("Olá Lucas!")
	default:
		fmt.Println("Olá estranho")
	}
}
