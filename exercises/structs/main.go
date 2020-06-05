package main

import "fmt"

type cannabis struct {
	name    string
	flowers []string
}

func main() {
	cS := cannabis{
		name: "Sativa",
	}
	cS = cS.updateFlower("Gorilla Haze")

	fmt.Println(cS)

}

func (c cannabis) updateFlower(nFlower string) cannabis {
	c.flowers = append(c.flowers, nFlower)
	return c
}
