package main

import "fmt"

func main() {
	type (
		cannabis []string
	)
	var (
		sativa cannabis
		indica cannabis
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

}
