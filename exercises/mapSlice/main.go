package main

import (
	"fmt"
	"os"
)

func main() {
	var (
		// a map doesn't contains the values it self, it just contains a pointer to map header

		// I could declare make(map[string]string) or map[string]string{}, both ways initialize the map
		ptEn = map[string]string{}
		args = os.Args[1:]
		flag bool
	)

	ptEn["Olá"], ptEn["Carro"] = "Hello", "Car"

	if len(args) == 1 {
		q := args[0]
	scope:
		for i, v := range ptEn {
			if q == i {
				fmt.Printf("The word %v is %v in english\n", q, v)
				flag = true
				break scope
			} else {
				continue
			}
		}
		switch flag {
		case false:
			fmt.Println("No word find")
		}
	} else {
		fmt.Println("Write one word")
	}
	// delete a value from a key in map
	delete(ptEn, "Olá")

	fmt.Println(ptEn)
}
