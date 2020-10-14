package main

import (
	"fmt"
	"time"
)

func main() {
	switch h := time.Now().Hour(); {
	case h >= 0 && h <= 12:
		fmt.Println("Good morning")
	case h >= 13 && h <= 18:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good Night")
	}
}
