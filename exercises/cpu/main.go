package main

import (
	f "fmt"
	"runtime"
)

func main() {
	f.Println(runtime.NumCPU())
}
