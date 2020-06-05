package main

import "fmt"

type (
	// Define flowers struct
	flowers struct {
		name string
		thc  int
		geo  map[string]int // declare a map inside the struct
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
		"Brasil": 2, // country and qtde of this flower at this country
	}
	// Add new sativa and indicas flowers
	cSativa, cIndica = cSativa.updateFlower("Gorilla Haze", 27, eX), cIndica.updateFlower("Notherland", 22, eX)

	// Printing the flowers for each type
	fmt.Println("\nThe Sativa flowers are:")
	for _, f := range cSativa.flowers {
		fmt.Printf("Name: %v ------- THC: %v\n", f.name, f.thc)
		for c, q := range f.geo {
			fmt.Printf("Country: %v ----- Qtde: %v\n\n", c, q)
		}
	}
	fmt.Println("----------------------------")
	fmt.Println("\n\nThe Indica flowers are:")

	for _, f := range cIndica.flowers {
		fmt.Printf("Name:%v ------- THC: %v\n", f.name, f.thc)
		for c, q := range f.geo {
			fmt.Printf("Country: %v ----- Qtde: %v\n\n", c, q)
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
