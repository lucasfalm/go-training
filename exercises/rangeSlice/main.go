package main

import (
	"fmt"
	"os"
)

func main() {
	for i, v := range os.Args[1:] {
		fmt.Printf("The index %v has the value %v\n", i + 1, v)
	}
}
