package main

import "fmt"

type (
	// Define flowers struct
	flowers struct {
		name string
		thc  int
		geo  map[string]int // declare a map inside a struct
	}
	// Define cannabis struct
	cannabis struct {
		name    string
		flowers []flowers
	}
)

func main() {
	// Create new types of cannabis using struct type
	cSativa, cIndica := cannabis{name: "Sativa"}, cannabis{name: "Indica"}

	// Example of geo
	eX := map[string]int{
		"Brasil": 2, // countr and qtde of this flower at this country
	}
	// Add new sativa and indicas flowers
	cSativa, cIndica = cSativa.updateFlower("Gorilla Haze", 27, eX), cIndica.updateFlower("Notherland", 22, eX)

	// Printing the flowers for each type
	fmt.Println("\nThe Sativa flowers are:")
	for _, f := range cSativa.flowers {
		fmt.Printf("Name: %v ------- THC: %v\n", f.name, f.thc)
		for i, v := range f.geo {
			fmt.Printf("Country: %v ----- Qtde: %v\n\n", i, v)
		}
	}
	fmt.Println("----------------------------")
	fmt.Println("\n\nThe Indica flowers are:")

	for _, g := range cIndica.flowers {
		fmt.Printf("Name:%v ------- THC: %v\n", g.name, g.thc)
		for i, v := range g.geo {
			fmt.Printf("Country: %v ----- Qtde: %v\n\n", i, v)
		}
	}
}

// Function to add new flowers for cannabis type
func (c cannabis) updateFlower(n string, thc int, g map[string]int) cannabis {
	nF := flowers{
		name: n,
		thc:  thc,
		geo:  make(map[string]int), // Initializing map
	}
	nF.geo = g                        // Give a value to the map
	c.flowers = append(c.flowers, nF) // Insert every info into flower
	return c
}
