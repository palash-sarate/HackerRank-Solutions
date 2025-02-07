package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "container/list"
)

/*
 * Complete the 'knightlOnAChessboard' function below.
 *
 * The function is expected to return a 2D_INTEGER_ARRAY.
 * The function accepts INTEGER n as parameter.
 */

// func isSliceEqual(ar1 []int32,ar2 []int32) bool{
//     if len(ar1) != len(ar2) {
//         return false
//     }
//     for i,v := range ar1 {
//         if(ar2[i]!=v){
//             return false;
//         }
//     }
//     return true
// }
type Position struct {
    x, y, moves int32
}
func getPossibleMoves(cPos []int32, a int32, b int32) [][]int32{
    var posMoves = [][]int32{
        {cPos[0]+a,cPos[1]+b},
        {cPos[0]+a,cPos[1]-b},
        {cPos[0]-a,cPos[1]+b},
        {cPos[0]-a,cPos[1]-b},
        {cPos[0]+b,cPos[1]+a},
        {cPos[0]+b,cPos[1]-a},
        {cPos[0]-b,cPos[1]+a},
        {cPos[0]-b,cPos[1]-a},
        };
    return posMoves;
}
func isValid(move []int32, boardSize int32) bool{
    return move[0] < boardSize && move[1] < boardSize && move[0] >= 0 && move[1] >= 0
}
func findMinStepsReq(sPos []int32, tPos []int32, a int32, b int32, boardSize int32) int32{
    // BFS init
    queue := list.New()
    queue.PushBack(Position{sPos[0], sPos[1], 0})
    visited := make(map[[2]int32]bool)
    visited[[2]int32{sPos[0], sPos[1]}] = true
    
    for queue.Len() > 0 {
        elem := queue.Front()
        queue.Remove(elem)
        pos := elem.Value.(Position)
        // if current position is target position return the moves required to reach there
        if (pos.x == tPos[0] && pos.y == tPos[1]){
            return int32(pos.moves);
        }
        // find possible moves for current position and add them to queue
        for _,move := range getPossibleMoves([]int32{pos.x,pos.y},a,b) {
            if(isValid(move, boardSize) && !visited[[2]int32{move[0],move[1]}]){
                visited[[2]int32{move[0],move[1]}] = true
                queue.PushBack(Position{move[0], move[1], pos.moves + 1})
                }
        }
    }
    return -1
}
func knightlOnAChessboard(n int32) [][]int32 {
    
    // Write your code here
    steps := make([][]int32, n-1)
    for i := range steps {
        steps[i] = make([]int32, n-1)
    }
    var targetPos = []int32{n-1,n-1};
    var startPos = []int32{0,0};
    for  b := int32(1); b < n; b++ {
        for  a := int32(1); a < n; a++ {
            // fmt.Println(a,b)
            steps[a-1][b-1] = findMinStepsReq(startPos, targetPos, a, b, n);
        }
    }
    return steps;
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

    result := knightlOnAChessboard(n)

    for i, rowItem := range result {
        for j, colItem := range rowItem {
            fmt.Fprintf(writer, "%d", colItem)

            if j != len(rowItem) - 1 {
                fmt.Fprintf(writer, " ")
            }
        }

        if i != len(result) - 1 {
            fmt.Fprintf(writer, "\n")
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
