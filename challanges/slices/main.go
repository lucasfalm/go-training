package main

import (
	"fmt"
	"os"
)

type (
	cannabis []string
	// multidimentional slice
	canabinoides [][]cannabis
)

func main() {

	var (
		sativa cannabis
		indica cannabis
		args   = os.Args[1:]
	)
	sativa = append(sativa, "Cannalope Haze", "Sour Kosher")
	indica = append(indica, "Afghan Kush", "Quick Kush")

	if len(args) >= 1 {
		cType := args[0]

		if cType == "sativa" {
			for _, f := range args[1:] {
				sativa.addCannabis(f)
			}
		} else if cType == "indica" {
			for _, f := range args[1:] {
				indica.addCannabis(f)
			}
		} else {
			fmt.Println("Choose [indica] [sativa]")
		}
	}

	fmt.Println("Sativas:")
	for _, v := range sativa {
		fmt.Printf("%v\n", v)
	}
	fmt.Println("\n\nIndicas:")
	for _, z := range indica {
		fmt.Printf("%v\n", z)
	}
}

func (c *cannabis) addCannabis(cName string) {
	*c = append(*c, cName)
}
