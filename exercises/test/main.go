package main

import (
	"fmt"

	libgiphy "github.com/sanzaru/go-giphy"
)

func main() {
	giphy := libgiphy.NewGiphy("9Jva56zlJfYORERbtbIbxPXEvq6PZzf0")

	dataRandom, err := giphy.GetRandom("very funny")
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("Random data: %+v\n", dataRandom)

}
