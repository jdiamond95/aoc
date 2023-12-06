package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadLines(p string) []string {
    f, err := os.Open(p)
    if err != nil {
        fmt.Println(err)
    }
    defer f.Close()

    sc := bufio.NewScanner(f)
    var lines []string

    for sc.Scan() {
        lines = append(lines, sc.Text())
    }

    return lines
}


func ReadBlock(p string) string {
    b, err := os.ReadFile(p)
    if err != nil {
        fmt.Println(err)
    }

    text := string(b)
    return text
}

func ReadIntArray(s string) []int {
    // Assume theres something colon splitting our line
    s = strings.Split(s, ":")[1]
    strArr := strings.Fields(s)

    intArr := make([]int, len(strArr))
    for i := range strArr {
        num, err := strconv.Atoi(strArr[i])
        if err != nil {
            fmt.Println(err)
        }
        
        intArr[i] = num
    }

    return intArr
}

func ReadIntKerning (s string) int {
    // Assumes theres a colon seperating line title and data
    numStr := strings.Split(s, ":")[1]
    numStr = strings.Replace(numStr, " ", "", -1)
    i, err := strconv.Atoi(numStr)

    if err != nil {
        fmt.Println(err)
    }
    
    return i
}
