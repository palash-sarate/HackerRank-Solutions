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
 * Complete the 'balancedSums' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */
// TODO: algo fails only when sums or elements are zero
// func balancedSums(arr []int32) string {
//     // Write your code here
//     n := len(arr)
//     var left, right int64 = -1, int64(n-1);
//     var sumleft, sumright int64 = 0, int64(arr[n-1]);
//     done := false;
//     var exist string = "NO";
    
//     // fmt.Println(arr,left,right,sumleft,sumright)
//     for done == false {
//         if(n==1){
//             done = true;exist = "YES"
//         }else if (left >= right || right - left == 1){
//             done = true;
//             exist = "NO"            
//         }else if (sumleft == sumright) {
//             // fmt.Println("equal sums")
//             if((right - left) == 2){
//                 done = true;
//                 exist = "YES"
//             }else{
//                 for sumright == (sumright + int64(arr[right-1])){right--}
//                 for sumleft == (sumleft + int64(arr[left+1])){left++}
//                 if((right - left) > 2){
//                     right--;
//                     sumright += int64(arr[right])
//                 }
//             }
//         } else if(sumleft > sumright) {
//             // fmt.Println("sumleft > sumright")
//             for sumleft > sumright{
//                 right--;sumright += int64(arr[right])
//             }
//         } else if(sumleft < sumright) { 
//             // fmt.Println("sumleft < sumright")
//             for sumleft < sumright {
//                 left++;
//                 sumleft += int64(arr[left])
//             }
//         }
//         // fmt.Println(left, right, sumleft, sumright)
//     }
//     // fmt.Println(exist)
//     return exist
// }
func balancedSums(arr []int32) string {
    totalSum := int32(0)
    for _, num := range arr {
        totalSum += num
    }

    leftSum := int32(0)
    for _, num := range arr {
        if leftSum == totalSum - leftSum - num {
            return "YES"
        }
        leftSum += num
    }

    return "NO"
}
func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    TTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    T := int32(TTemp)

    for TItr := 0; TItr < int(T); TItr++ {
        nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        n := int32(nTemp)

        arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

        var arr []int32

        for i := 0; i < int(n); i++ {
            arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
            checkError(err)
            arrItem := int32(arrItemTemp)
            arr = append(arr, arrItem)
        }

        result := balancedSums(arr)

        fmt.Fprintf(writer, "%s\n", result)
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
 * Complete the 'balancedSums' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */
// TODO: algo fails only when sums or elements are zero
func balancedSums(arr []int32) string {
    // Write your code here
    n := len(arr)
    var left, right int64 = -1, int64(n-1);
    var sumleft, sumright int64 = 0, int64(arr[n-1]);
    done := false;
    var exist string = "NO";
    
    // fmt.Println(arr,left,right,sumleft,sumright)
    for done == false {
        if(n==1){
            done = true;exist = "YES"
        }else if (left >= right || right - left == 1){
            done = true;
            exist = "NO"            
        }else if (sumleft == sumright) {
            // fmt.Println("equal sums")
            if((right - left) == 2){
                done = true;
                exist = "YES"
            }else{
                for sumright == (sumright + int64(arr[right-1])){right--}
                for sumleft == (sumleft + int64(arr[left+1])){left++}
                if((right - left) > 2){
                    right--;
                    sumright += int64(arr[right])
                }
            }
        } else if(sumleft > sumright) {
            // fmt.Println("sumleft > sumright")
            for sumleft > sumright{
                right--;sumright += int64(arr[right])
            }
        } else if(sumleft < sumright) { 
            // fmt.Println("sumleft < sumright")
            for sumleft < sumright {
                left++;
                sumleft += int64(arr[left])
            }
        }
        // fmt.Println(left, right, sumleft, sumright)
    }
    // fmt.Println(exist)
    return exist
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    TTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    T := int32(TTemp)

    for TItr := 0; TItr < int(T); TItr++ {
        nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
        checkError(err)
        n := int32(nTemp)

        arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

        var arr []int32

        for i := 0; i < int(n); i++ {
            arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
            checkError(err)
            arrItem := int32(arrItemTemp)
            arr = append(arr, arrItem)
        }

        result := balancedSums(arr)

        fmt.Fprintf(writer, "%s\n", result)
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
