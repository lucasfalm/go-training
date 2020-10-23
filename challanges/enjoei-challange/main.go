package main

import (
	"fmt"
	"sort"
)

/*
2 parametros - 1 parametro é um array de inteiros (grande), segundo parametro é um int
bool como resposta - soma de 2 numeros é igual ao segundo paramtro

-------
examples=

1 input = [0,50,-25,2,10,3,15,11,1]
2 input = 5

------

inteiro pode ser negativo
*/

// [0,10, 5,11], 5
// [0, 5, 10, 11], 5
func main() {
	fmt.Println(challenge([]int{0, 50, -25, 2, 10, 3, 15, 11, 1}, 13))
}

// Big O(n log n)
func challenge(arr []int, sum int) bool {
	var (
		pointerOne int
		pointerTwo = len(arr) - 1
	)

	sort.Ints(arr)

	for pointerOne < pointerTwo {
		actualSum := arr[pointerOne] + arr[pointerTwo]

		if actualSum == sum {
			return true
		} else if actualSum < sum {
			pointerOne++
		} else {
			pointerTwo--
		}
	}
	return false
}
