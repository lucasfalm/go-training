package main

import (
	"fmt"
	"os"
)

type (
	cannabis []string
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
			for _, v := range args[1:] {
				sativa = sativa.addCannabis(v)
			}
		} else if cType == "indica" {
			for _, v := range args[1:] {
				indica = indica.addCannabis(v)
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

func (c cannabis) addCannabis(cName string) cannabis {
	c = append(c, cName)
	return c
}
