package main

import (
	"fmt"
	"strconv"
)

type (
	product struct {
		name  string
		value interface{}
	}
)

func (p *product) print() {
	fmt.Printf("%v ------ cost: U$%v\n", p.name, p.value)
}

func (p *product) discount(d int) {
	switch p.value.(type) {
	case int:
		p.value = applyDiscount(p.value.(int), d)
		printDiscount(p.name, p.value.(int))
	case string:
		cV, _ := strconv.Atoi(p.value.(string))
		p.value = applyDiscount(cV, d)
		printDiscount(p.name, p.value.(int))
	default:
		fmt.Println("Error")
	}
}

func applyDiscount(aP int, d int) int {
	return aP - ((aP * d) / 100)
}

func printDiscount(name string, value int) {
	fmt.Printf("\nNew price of %v ------ cost: U$%v\n\n", name, value)
}
