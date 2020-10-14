package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'pickingNumbers' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */
// https://www.hackerrank.com/challenges/picking-numbers/problem
func pickingNumbers(a []int32) int32 {
	rank := make(map[int]int)

	for _, value := range a {
		rank[int(value)]++
	}

	keys := make(map[int]int)

	for n, qtde := range rank {
		keys[qtde] = n
	}

	fmt.Println(keys)
	fmt.Println(rank)
	first := 0
	firstQtde := 0

	for qtde, value := range keys {
		if qtde > firstQtde {
			first = value
		}

		fmt.Println(qtde, value)
	}

	counter := 0
	const max = 10
	flag := false
	firstCopie := first
	fmt.Println(rank[first])

	for counter > max {
		if !flag {
			if _, ok := rank[first+1]; ok {
				return int32(rank[first] + rank[first+1])
			}

			first++
			counter++

			if counter > max {
				flag = true
				counter = 0
				first = firstCopie
			}
		} else {
			if _, ok := rank[first-1]; ok {
				return int32(rank[first] + rank[first-1])
			}
			first--
		}
	}
	return 0
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	result := pickingNumbers(a)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
