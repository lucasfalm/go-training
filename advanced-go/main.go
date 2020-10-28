package main

import "fmt"

func main() {
	one, two, three := myfunc(1, 2, 3)
	fmt.Println(one, two, three)

	switch one {
	case 1002:
		fmt.Println("1002")
		fallthrough
	default:
		fmt.Println("will also pass here")
	}

}

func myfunc(a, b, c int) (one, two, three int) {
	one = a + 1
	changingValues(&one, 1000)

	two = one + 2
	three = *(notafunc())
	return // the returned values were defined explicity
}

func notafunc() (number *int) {
	i := 55
	number = &(i)
	return
}

func changingValues(value *int, newValue int) {
	*value = newValue + *value
	return
}
