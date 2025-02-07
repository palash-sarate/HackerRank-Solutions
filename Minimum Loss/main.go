package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "math"
    "sort"
)

/*
 * Complete the 'minimumLoss' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts LONG_INTEGER_ARRAY price as parameter.
 */

func minimumLoss(prices []int64) int32 {
    type priceItem struct{
        price int64 
        index int32 
    }
    priceList := make([]priceItem,len(prices));
    for i,v := range prices{
        priceList[i] = priceItem{v,int32(i)};
    }
    // sort priceList in ascending order
    sort.Slice(priceList, func(i, j int) bool {
        return priceList[i].price < priceList[j].price
    })
    // Write your code here
    var minLoss int64 = math.MaxInt64;
    for i:= 0;i<len(priceList)-1;i++ {
        if(priceList[i].index > priceList[i+1].index){
            loss := priceList[i+1].price - priceList[i].price;
            fmt.Println(loss)
            if(minLoss > int64(loss)){minLoss = int64(loss)}
        }        
    }
    return int32(minLoss);
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

    priceTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var price []int64

    for i := 0; i < int(n); i++ {
        priceItem, err := strconv.ParseInt(priceTemp[i], 10, 64)
        checkError(err)
        price = append(price, priceItem)
    }

    result := minimumLoss(price)

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
