package main

import "fmt"

type (
	// Define flowers struct
	flowers struct {
		name string
		thc  int
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

	// Add new sativa and indicas flowers
	cSativa = cSativa.updateFlower("Gorilla Haze", 27)
	cIndica = cIndica.updateFlower("Notherland", 22)

	// Printing the flowers for each type
	fmt.Println("\nThe Sativa flowers are:")
	for _, f := range cSativa.flowers {
		fmt.Printf("Name: %v ------- THC: %v\n", f.name, f.thc)
	}

	fmt.Println("\n\nThe Indica flowers are:")
	for _, g := range cIndica.flowers {
		fmt.Printf("Name:%v ------- THC: %v\n", g.name, g.thc)
	}

}

// Function to add new flowers for cannabis type
func (c cannabis) updateFlower(n string, thc int) cannabis {
	nF := flowers{
		name: n,
		thc:  thc,
	}
	c.flowers = append(c.flowers, nF)
	return c
}
