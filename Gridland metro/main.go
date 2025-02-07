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
 * Complete the 'gridlandMetro' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER m
 *  3. INTEGER k
 *  4. 2D_INTEGER_ARRAY track
 */

func gridlandMetro(n int64, m int64, k int64, track [][]int64) int64 {
    // Sort by the second element (start of track)
    sort.Slice(track, func(i, j int) bool {
        return track[i][1] < track[j][1] // Sort in ascending order
    })
    var lamposts int64 = n * m;
    var tracks = make(map[int64][]int64);
    // Write your code here
    for _,v := range track {
        if old_track, exists := tracks[v[0]]; exists {
            // if track exists in tracks
            // new track can be disjoint or merged
            // if disjoint ie. old end < new start
            if(old_track[1] < v[1]){
                // reduce new track's length from lamposts
                lamposts -= v[2]-v[1]+1;
                // update old track's end to new track end
                tracks[v[0]][1] = v[2];
            }
            // if merged old end >= new start and old end < new end
            if(old_track[1] >= v[1] && old_track[1] < v[2]){
                // new track start = old track end
                v[1] = old_track[1] + 1;
                // reduce new track's length from lamposts
                lamposts -= v[2]-v[1]+1;
                // update old track's end to new track end
                tracks[v[0]][1] = v[2];
            }
        } else {
            // Key i does not exist
            tracks[v[0]] = []int64{v[1],v[2]};
            lamposts -= v[2]-v[1]+1;            
        }
        
        fmt.Println(v, lamposts)
    }    
    return lamposts
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
    n := int64(nTemp)

    mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
    checkError(err)
    m := int64(mTemp)

    kTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
    checkError(err)
    k := int64(kTemp)

    var track [][]int64
    for i := 0; i < int(k); i++ {
        trackRowTemp := strings.Split(strings.TrimRight(readLine(reader)," \t\r\n"), " ")

        var trackRow []int64
        for _, trackRowItem := range trackRowTemp {
            trackItemTemp, err := strconv.ParseInt(trackRowItem, 10, 64)
            checkError(err)
            trackItem := int64(trackItemTemp)
            trackRow = append(trackRow, trackItem)
        }

        if len(trackRow) != 3 {
            panic("Bad input")
        }

        track = append(track, trackRow)
    }

    result := gridlandMetro(n, m, k, track)

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
