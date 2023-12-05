package utils

import (
	"bufio"
	"fmt"
	"os"
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
