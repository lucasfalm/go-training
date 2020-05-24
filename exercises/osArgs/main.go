package main

import (
	"fmt"
	"os"
)

func main() {
	var args []string
	args = os.Args
	/*
		fmt.Println("Path", args[0])
		fmt.Println("1th", args[1])
		fmt.Println("2th", args[2])
		fmt.Println("3th", args[3])
	*/

	for i := 0; i < len(args); i++ {
		fmt.Println("Argument position", i, ":", args[i])
	}

	fmt.Printf("%#v\n", args)
	fmt.Println("Number of itens at slice:", len(args))
}
