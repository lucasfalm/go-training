package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// https://www.hackerrank.com/challenges/beautiful-days-at-the-movies/problem
// Complete the beautifulDays function below.
func beautifulDays(i int32, j int32, k int32) int32 {
	var (
		total int32
	)

	for i <= j {
		var (
			sliceActualDay    = strings.Split(strconv.Itoa(int(i)), "")
			counter           = len(sliceActualDay) - 1
			stringReversedDay string
			rest              int
			reversedDay       int
		)

		for counter >= 0 {
			stringReversedDay += sliceActualDay[counter]
			counter--
		}

		reversedDay, _ = strconv.Atoi(stringReversedDay)
		rest = reversedDay - int(i)

		if rest%int(k) == 0 {
			total++
		}
		i++
	}
	return total
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	ijk := strings.Split(readLine(reader), " ")

	iTemp, err := strconv.ParseInt(ijk[0], 10, 64)
	checkError(err)
	i := int32(iTemp)

	jTemp, err := strconv.ParseInt(ijk[1], 10, 64)
	checkError(err)
	j := int32(jTemp)

	kTemp, err := strconv.ParseInt(ijk[2], 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := beautifulDays(i, j, k)

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
