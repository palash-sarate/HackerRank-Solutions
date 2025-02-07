package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "sort"
)

/*
 * Complete the 'pairs' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY arr
 */
 
//  func sortAsc(arr []int32) []int32{
//     sort.Slice(arr, func(i, j int) bool {
//         return arr[i] < arr[j]
//     })
//     return arr
// }
func pairs(k int32, arr []int32) int32 {
    // Write your code here
    sort.Slice(arr, func(i, j int) bool {
        return arr[i] < arr[j]
    })
    // fmt.Println(arr, k)
    var pairs int32 = 0;
    // brute force
    // for i := 1; i < len(arr); i++ {
    //     for j := 0; j<i; j++ {  
    //         // fmt.Println(arr[i], arr[j],arr[i] - arr[j], arr[i] - arr[j] == k)          
    //         if (arr[i] - arr[j] == k){
    //             pairs ++;
    //             break
    //         }
    //     }
    // }
    
    // two pointer method
    left, right := 0, 1
    n := len(arr)
    for right < n {
        diff := arr[right] - arr[left]
        if (diff == k){
            pairs++
            left++
            right++
        } else if(diff < k){
            right++
        } else{ // diff > k
            left++
            if left == right{
                right++
            }
        }
    }
    return pairs
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
    checkError(err)
    k := int32(kTemp)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    result := pairs(k, arr)

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
