package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// https://www.hackerrank.com/challenges/sock-merchant/problem
// Complete the sockMerchant function below.
// Big O(n) time and O(n) space complexity 
func sockMerchant(n int32, ar []int32) int32 {
    couples := map[int]int{}
    counter := 0
    result := 0

    for counter <= len(ar) - 1 {
        couples[int(ar[counter])]++
        counter++
    }
    
    for _, qtde := range couples {
        if qtde % 2 == 0 {
            result += qtde / 2
        } else {
            if qtde - 1 >= 2 {
                qtde--
                result += qtde / 2
            }
        } 
    }
    return int32(result)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arTemp := strings.Split(readLine(reader), " ")

    var ar []int32

    for i := 0; i < int(n); i++ {
        arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
        checkError(err)
        arItem := int32(arItemTemp)
        ar = append(ar, arItem)
    }

    result := sockMerchant(n, ar)

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

