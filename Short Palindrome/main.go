package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

/*
 * Complete the 'shortPalindrome' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */
func isPalindrome(s string) bool {
    runes := []rune(s) // Convert to rune slice to handle Unicode correctly
    n := len(runes)

    for i, v := range runes {
        if v != runes[n-i-1] { // Compare from start and end
            return false
        }
    }
    return true
}

func shortPalindrome(s string) int32 {
    s_len := len(s)
    // Write your code here
    n := 4
    pointers := make([]int, n) // Create a slice of length n
	counter := int32(0);
    for i := 0; i < n; i++ {
        pointers[i] = i // Assign consecutive values
    }
    
    for pointers[0] == s_len - n + 1{
        
    } 
return counter
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    result := shortPalindrome(s)

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
