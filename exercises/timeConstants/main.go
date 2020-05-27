package main

import (
	"fmt"
)

func main() {
	const (
		EST = -(5 + iota)
		_
		MST
		PST
	)

	fmt.Printf("The EST time is %v\n and the MST time is %v\n and the PST time is %v\n", EST, MST, PST)
}
