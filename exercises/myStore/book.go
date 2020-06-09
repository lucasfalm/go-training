package main

import "fmt"

type (
	book struct {
		name  string
		value interface{}
	}
)

func (b book) print() {
	fmt.Printf("%v ------ cost: U$%v\n", b.name, b.value)
}
