package main

import (
	"fmt"
	"os"
)

func main() {

	var (
		args          = os.Args
		validUser     = "flucas97"
		validPassword = "123456"
		userName      string
		password      string
	)

	const (
		usage    = "Usage: [username] [password]"
		denied   = "Access denied for %v\n"
		errPwd   = "Invalid password for %v\n"
		accessOK = "Access granted to %v\n"
	)

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	if userName, password = os.Args[1], os.Args[2]; userName != validUser {
		fmt.Printf(denied, userName)
	} else if password != validPassword {
		fmt.Printf(errPwd, userName)
	} else {
		fmt.Printf(accessOK, userName)
	}
}
