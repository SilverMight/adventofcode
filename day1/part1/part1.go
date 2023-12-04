package main

import (
	"bufio"
	"fmt"
	"os"
)

func getFirstDigit(line string) (uint) {

    for _, char := range line {
        if num := uint(char - '0') ; num >= 0 && num <= 9 {
            return num
        }
    }
    return 0
}

func getLastDigit(line string) (uint) {

    for i := len(line) - 1; i >= 0; i-- {
        if num := uint(line[i] - '0') ; num >= 0 && num <= 9 {
            return num
        }
    }
    return 0
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    var sum uint 

    for scanner.Scan() {
        line := scanner.Text()
       
        sum += getFirstDigit(line) * 10 + getLastDigit(line)

    }
    fmt.Println(sum)
}
