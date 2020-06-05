package main

import "fmt"

func main() {
	var (
		letters [][]byte
		names   []string
	)

	lucas, maiara := "lucas", "maiara"
	bL, bM := []byte(lucas), []byte(maiara)

	letters = append(letters, bL, bM)
	names = append(names, lucas, maiara)

	for i := 0; i < len(letters); i++ {
		fmt.Printf("The name %v is %v in bytes\n", names[i], letters[i])
	}

}
