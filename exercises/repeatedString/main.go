package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://www.hackerrank.com/challenges/repeated-string/problem
// Complete the repeatedString function below.
// Big O(n) time( where n is length of string) and Big O(n) space. 16/23 tests
func repeatedString(s string, n int64) int64 {
	var (
		lettersSlice = strings.Split(s, "")
		repeat       int
	)

	for _, letter := range lettersSlice {
		if letter == "a" {
			repeat++
		}
	}

	return int64(math.Round(float64(repeat) / (float64(len(lettersSlice))) * float64(n)))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	n, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	result := repeatedString(s, n)

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
