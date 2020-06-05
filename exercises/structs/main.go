package main

import "fmt"

type flowers struct {
	name string
	thc  int
}
type cannabis struct {
	name    string
	flowers []flowers
}

func main() {
	cSativa := cannabis{
		name: "Sativa",
	}
	cSativa = cSativa.updateFlower("Gorilla Haze", 27)

	cIndica := cannabis{
		name: "Indica",
	}

	cIndica = cIndica.updateFlower("Notherland", 22)

	fmt.Println("\nThe Sativa flowers are:")
	for _, f := range cSativa.flowers {
		fmt.Printf("Name: %v ------- THC: %v\n", f.name, f.thc)
	}

	fmt.Println("\n\nThe Indica flowers are:")
	for _, f := range cIndica.flowers {
		fmt.Printf("Name:%v ------- THC: %v\n", f.name, f.thc)
	}

}

func (c cannabis) updateFlower(n string, thc int) cannabis {
	nF := flowers{
		name: n,
		thc:  thc,
	}
	c.flowers = append(c.flowers, nF)
	return c
}
