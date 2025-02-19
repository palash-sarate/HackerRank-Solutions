package main

import (
    // "bufio"
    "fmt"
    // "io"
    // "os"
    // "strconv"
    // "strings"
)

/*
 * Complete the 'connectedCell' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY matrix as parameter.
 */
type loc struct {
    A, B int32
}

func connectedCell(matrix [][]int32) int32 {
    visited := make(map[loc]bool)
    // Write your code here
    var findConnected func(i int32, j int32) int32
    var maxArea int32 = 0;
    var width = int32(len(matrix));
    var height = int32(len(matrix[0]));
    
    directions := []struct{ di, dj int32 }{
        {-1, -1}, {-1, 0}, {-1, 1},
        { 0, -1},          { 0, 1},
        { 1, -1}, { 1, 0}, { 1, 1},
    }
	findConnected = func(l int32, m int32) int32 {
        if l < 0 || m < 0 || l >= width || m >= height || matrix[l][m] == 0 || visited[loc{l, m}] {
            return 0
        }
        visited[loc{l,m}] = true;
		count := int32(1)
        for _, d := range directions {
            count += findConnected(l+d.di, m+d.dj)
        }
        return count
    }
    for i,row := range matrix {
        for j,val := range row {
            var area int32 = 1;
            if (visited[loc{int32(i),int32(j)}] || val == 0){continue}
            area = findConnected(int32(i),int32(j));
            if area > maxArea {maxArea = area;}
        }
    }

    return maxArea
}

func main() {
    input_matrix := [][]int32{
        // {1, 1, 0, 0},
        // {0, 1, 1, 0},
        // {0, 0, 1, 0},
        // {1, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 1, 1, 0},
		{1, 1, 0, 0, 1, 0, 0, 0, 1},
		{0, 0, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 1, 1, 0, 1, 0, 1, 1},
		{0, 1, 1, 1, 0, 0, 1, 1, 0},
		{0, 1, 0, 1, 1, 0, 1, 1, 0},
		{0, 1, 0, 0, 1, 1, 0, 1, 1},
		{1, 0, 1, 1, 1, 1, 0, 0, 0},
    }

    fmt.Println(connectedCell(input_matrix))
}


