package main

import "fmt"

type (
	product struct {
		name  string
		value interface{}
	}
)

func (p *product) print() {
	fmt.Printf("%v ------ cost: U$%v\n", p.name, p.value)
}
