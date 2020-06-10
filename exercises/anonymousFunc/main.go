package main

import "fmt"

func main() {
	func(h string) {
		fmt.Println(h)
	}("Hello world!")
}
