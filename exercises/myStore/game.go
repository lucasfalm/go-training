package main

import "fmt"

type game struct {
	name  string
	value interface{}
}

func (g game) print() {
	fmt.Printf("%v ------ cost: U$%v\n", g.name, g.value)
}
