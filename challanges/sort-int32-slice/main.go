// https://www.hackerrank.com/challenges/countingsort2/problem
func countingSort(arr []int32) []int32 {
    intSlice := []int{}
    result := []int32{}

    for _, number := range arr { 
        intSlice = append(intSlice, int(number))
    }
    sort.Ints(intSlice)

    for _, number := range intSlice { 
        result = append(result, int32(number))
    }
    return result
}
