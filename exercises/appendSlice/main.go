package main

import (
	"fmt"
	"os"
)

var (
	flag    bool
	unique  string
	args    = os.Args[1:]
	numbers = make(map[int]string)
	uniques []string
)

func main() {
	if len(args) >= 1 {
		for i, value := range args {
			flag = false
			numbers[i] = value

			if len(uniques) > 0 {
				for _, number := range uniques {
					if number == value {
						flag = true
					}
				}
			}
			if flag == false {
				uniques = append(uniques, value)
			}
		}
	} else {
		fmt.Println("Write some numbers!")
		return
	}

	fmt.Println("All numbers:")
	for i := 0; i <= len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	fmt.Printf("Size: %v\n", len(numbers))

	fmt.Println("--------------------")

	fmt.Println("All unique numbers:")

	for _, unique := range uniques {
		fmt.Println(unique)
	}
	fmt.Printf("Size: %v\n", len(uniques))
}
