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

	fmt.Println("Sativas:")
	for _, v := range sativa {
		fmt.Printf("%v\n", v)
	}
	fmt.Println("\n\nIndicas:")
	for _, z := range indica {
		fmt.Printf("%v\n", z)
	}
	if len(args) >= 1 {
		sativa = sativa.addCannabis(args[0])
	}

	fmt.Println(sativa)
}

func (c cannabis) addCannabis(cName string) cannabis {
	c = append(c, cName)
	return c
}
