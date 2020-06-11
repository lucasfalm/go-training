package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo() // if defer is defined before err occurs
	f, err := os.Open("no-file.txt")
	if err != nil {
		log.Panic(err) // or just panic(err)
	}
	defer f.Close()

	log.SetOutput(f) // log what happend
}

func foo() {
	fmt.Println("deffer is called if is panic")
}
