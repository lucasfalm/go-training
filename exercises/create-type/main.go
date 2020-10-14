package main

import (
	"fmt"

	"github.com/flucas97/exercises/createType/address"
)

func main() {
	type phoneType int64

	var phone phoneType

	phone = phoneType(32323232)

	fmt.Printf("My phone is %v and the type is: %T\n", phone, phone)

	var myAddress address.Address

	myAddress = "Florianópolis"

	fmt.Printf("My address is %v and the type is %T\n", myAddress, myAddress)
}
