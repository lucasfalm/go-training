package main

import "fmt"

type mT struct {
	num int
}

func main() {
	var (
		counter byte = 10
	)

	p := &counter
	v := *p

	fmt.Printf("The variable p stores the memory address %v, and its value is %v.\n", p, v)
	fmt.Printf("Its type is %T, and its value is of type %T.\n", p, v)

	g := mT{num: 1}
	fmt.Println(g.num)
	g.addOne()
	fmt.Println(g.num)
	fmt.Printf("Original g memory position: %p\n", &g)
}

func (c *mT) addOne() {
	c.num++
	fmt.Printf("Memory position %p\n", &c)
	fmt.Printf("Points to %p\n", c)
}
