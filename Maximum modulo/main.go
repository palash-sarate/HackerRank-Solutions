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
 * Complete the 'maximumSum' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts following parameters:
 *  1. LONG_INTEGER_ARRAY a
 *  2. LONG_INTEGER m
 */
func sumArr (a []int64) int64{
    var sum = int64(0)
    for _,v := range a {
        sum += v;
    }
    return sum
}
// func maximumSum(a []int64, m int64) int64 {
//     // Write your code here
//     n := len(a)
//     maxModulo := int64(0);
//     for i:=1;i<n+1;i++{
//         var sum int64 = 0;
//         for j:=0;j<n;j++ {
//            if(j+i <= n) {
//                 if(j == 0){
//                     sum = sumArr(a[j:j+i]);
//                 }else{
//                     sum = sum - a[j-1] + a[j+i-1];
//                 }
//                 modulo := sum%m
//                 if (maxModulo < modulo){maxModulo = modulo}
//                 if(maxModulo == m-1){return maxModulo}
//            }
//         }
//     }
//     return maxModulo
// }

// This Algo removes the recalculations of the sum of the subarray by using a prefix sum array
func maximumSum(a []int64, m int64) int64 {
    // Write your code here
    n := len(a)
    maxModulo := int64(0);
    // make a prefix sum array
    prefixSum := make([]int64, n)
    prefixSum[0] = a[0]
    for i:=1;i<n;i++{
        prefixSum[i] = prefixSum[i-1] + a[i]
    }
    
    for i:=0;i<n;i++{
        for j:=i;j<n;j++{
            var sum int64 = 0;
            if(i == 0){
                sum = prefixSum[j]
            }else{
                sum = prefixSum[j] - prefixSum[i-1]
            }
            modulo := sum%m
            if (maxModulo < modulo){maxModulo = modulo}
            if(maxModulo == m-1){return maxModulo}
        }
    }

    return maxModulo
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    q := int32(qTemp)

    for qItr := 0; qItr < int(q); qItr++ {
        firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

        nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
        checkError(err)
        n := int32(nTemp)

        m, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
        checkError(err)

        aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

        var a []int64

        for i := 0; i < int(n); i++ {
            aItem, err := strconv.ParseInt(aTemp[i], 10, 64)
            checkError(err)
            a = append(a, aItem)
        }

        result := maximumSum(a, m)

        fmt.Fprintf(writer, "%d\n", result)
    }

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
