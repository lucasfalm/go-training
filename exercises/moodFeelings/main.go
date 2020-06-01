package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var (
		args = os.Args[1:]
		name string
	)

	const (
		happy  = "feels happy"
		hungry = "feels hungry"
		mad    = "feels mad"
		sad    = "feels sad"
	)

	//mood := [...]string{happy, hungry, mad, sad}
	type moods [4]string

	mood := moods{happy, hungry, mad, sad}

	if len(args) == 1 {
		name = args[0]

		rand.Seed(time.Now().UnixNano())
		sNum := rand.Intn(len(mood))

		fmt.Printf("%s %v\n", name, mood[sNum])

	} else {
		fmt.Println("[Please, write a name]")
	}
}
