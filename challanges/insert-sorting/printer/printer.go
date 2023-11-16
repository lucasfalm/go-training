package printer

import (
	"fmt"
	"strconv"
	"strings"
)

func PrintFormatted(arr []int32) {
	var f = make([]string, len(arr))

	for i, n := range arr {
		f[i] = strconv.Itoa(int(n))
	}

	fmt.Println(strings.Join(f, " "))
}
