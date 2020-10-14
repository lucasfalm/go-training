package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// https://www.hackerrank.com/challenges/bon-appetit/problem
// Complete the bonAppetit function below.
// Big O(n) time and O(1) space complexity
func bonAppetit(bill []int32, k int32, b int32) {
    counter, total_bill, brian_total, anna_total := 0, 0, 0, 0 
    
    for counter <= len(bill) - 1 {
        total_bill += int(bill[counter])
        counter += 1
    }

    anna_total = int(bill[k])
    brian_total = (total_bill - anna_total) / 2
    overcharged := (brian_total + int(b) + anna_total) - total_bill

    switch overcharged { 
        case 0:
            fmt.Println("Bon Appetit")
            return
        default:
            fmt.Println(overcharged)
    }
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nk := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    nTemp, err := strconv.ParseInt(nk[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    kTemp, err := strconv.ParseInt(nk[1], 10, 64)
    checkError(err)
    k := int32(kTemp)

    billTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var bill []int32

    for i := 0; i < int(n); i++ {
        billItemTemp, err := strconv.ParseInt(billTemp[i], 10, 64)
        checkError(err)
        billItem := int32(billItemTemp)
        bill = append(bill, billItem)
    }

    bTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    b := int32(bTemp)

    bonAppetit(bill, k, b)
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

