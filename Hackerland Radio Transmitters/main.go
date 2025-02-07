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
 * Complete the 'hackerlandRadioTransmitters' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY x
 *  2. INTEGER k
 */

func hackerlandRadioTransmitters(x []int32, k int32) int32 {
    // Write your code here
    // x is sorted asc
    var antennas int32 = 0;
    sort.Slice(x[:],func(i,j int) bool{
        return x[i] < x[j]
    })
    fmt.Println(x)
    for i := 0; i<len(x); i++  {
        fmt.Println("i: ",i);
        // left side
        last_house_idx1 := findLastHouseUnderRange(x,k,int32(i));
        fmt.Println("right most for:",i,"is",last_house_idx1);
        // right side
        last_house_idx2 := findLastHouseUnderRange(x,k,last_house_idx1);
        fmt.Println("right most for:",last_house_idx1,"is",last_house_idx2);
        
        antennas += 1;
        // fmt.Println(last_house_idx,antennas);
        i = int(last_house_idx2);
        fmt.Println("next i: ",i+1);        
    }
    return antennas;
}

func findLastHouseUnderRange(x []int32, rangee int32, hidx int32) int32 {
    var fidx int32 = hidx
    for i := hidx; i<int32(len(x)); i++ {
        fmt.Println(x[i] - x[hidx], hidx, i);
        if((x[i] - x[hidx]) > rangee){
            return i-1;
        }
        fidx = i
    }
    // if no other house under the range of hidx it also becomes the fidx
    return fidx;
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

    xTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var x []int32

    for i := 0; i < int(n); i++ {
        xItemTemp, err := strconv.ParseInt(xTemp[i], 10, 64)
        checkError(err)
        xItem := int32(xItemTemp)
        x = append(x, xItem)
    }

    result := hackerlandRadioTransmitters(x, k)

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
