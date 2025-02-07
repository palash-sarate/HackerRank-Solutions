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
 * Complete the 'missingNumbers' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY arr
 *  2. INTEGER_ARRAY brr - reference 2
 */
func sortAsc(arr []int32) []int32{
    sort.Slice(arr, func(i, j int) bool {
        return arr[i] < arr[j]
    })
    return arr
}
func missingNumbers(arr []int32, brr []int32) []int32 {
    // Write your code here
    maxlength := len(brr);    
    if(len(arr) > maxlength){maxlength = len(arr)}
    
    intmap1 := make(map[int32]int32);
    intmap2 := make(map[int32]int32);
    var missingNum []int32;
        
    for i := 0; i<maxlength;i++{
        if i < len(arr) { 
            intmap1[arr[i]]++
        }
        if i < len(brr) { 
            intmap2[brr[i]]++
        }
    }
    
    // loop over intmap2
    for key, occ2 := range intmap2 {
        // check if v exists in arr
        if occ1, exists := intmap1[key]; exists {
            // fmt.Printf("Key %d exists with value: %s\n", key, value)
            if(occ1 != occ2){
                missingNum = append(missingNum, key);                
            }
        } else {
            // fmt.Printf("Key %d does not exist\n", key)
            missingNum = append(missingNum, key);
        }
    }
    return sortAsc(missingNum)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

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

    mTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    m := int32(mTemp)

    brrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var brr []int32

    for i := 0; i < int(m); i++ {
        brrItemTemp, err := strconv.ParseInt(brrTemp[i], 10, 64)
        checkError(err)
        brrItem := int32(brrItemTemp)
        brr = append(brr, brrItem)
    }

    result := missingNumbers(arr, brr)

    for i, resultItem := range result {
        fmt.Fprintf(writer, "%d", resultItem)

        if i != len(result) - 1 {
            fmt.Fprintf(writer, " ")
        }
    }

    fmt.Fprintf(writer, "\n")

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
