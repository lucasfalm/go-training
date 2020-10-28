package main

import "fmt"

func main() {
	one, two, three := myfunc(1, 2, 3)
	fmt.Println(one, two, three)
}

func myfunc(a, b, c int) (one, two, three int) {
	one = a + 1
	changingValues(&one, 1000)

	two = one + 2
	three = *(notafunc())
	return // the returned values were defined explicity
}

func notafunc() *int {
	number := 55
	return &number
}

func changingValues(value *int, newValue int) {
	*value = newValue + *value
	return
}
