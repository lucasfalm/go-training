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
	// verify what is the type of p.value (string or int)
	switch p.value.(type) {
	case int:
		p.value = applyDiscount(p.value.(int), d)
		printDiscount(p.name, p.value.(int))
	case string:
		actualPrice, _ := strconv.Atoi(p.value.(string))
		p.value = applyDiscount(actualPrice, d)
		printDiscount(p.name, p.value.(int))
	default:
		fmt.Println("Error")
	}
}

func applyDiscount(actualPrice int, d int) int {
	return actualPrice - ((actualPrice * d) / 100)
}

func printDiscount(name string, value int) {
	fmt.Printf("\nNew price of %v ------ cost: U$%v\n\n", name, value)
}
