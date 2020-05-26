package main

import (
	"fmt"
)

func main() {
	type phoneType int64

	var phone phoneType = 32323232

	fmt.Printf("My phone is %v and the type is: %T\n", phone, phone)
}
