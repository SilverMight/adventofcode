package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func isSymbol(char rune) bool {
    // Period
    if(char == '.') {
        return false
    }
    // Digit
    if(unicode.IsDigit(char)) {
        return false
    }

    return true
}

func hasAdjacentSymbol(lineIndex, charIndex, startIndex, endIndex int, arrayOfLines []string) bool {

    fmt.Println("Checking")

    if startIndex > 0 {
        if isSymbol(rune(arrayOfLines[lineIndex][startIndex - 1])) {
            return true
        }
    }

    // To the right of us
    if endIndex < len(arrayOfLines[lineIndex]) {
        if isSymbol(rune(arrayOfLines[lineIndex][endIndex])) {
            return true
        }
    }

    fmt.Println("bottom right")
    // bottom right diagonal
    if endIndex < len(arrayOfLines[lineIndex]) - 1 && lineIndex < len(arrayOfLines) - 1 {
        if isSymbol(rune(arrayOfLines[lineIndex + 1][endIndex])) {
            return true
        }
    }
    fmt.Println("top left")
    // top left diagonal 
    if startIndex > 0 && lineIndex > 0 {
        if isSymbol(rune(arrayOfLines[lineIndex - 1][startIndex - 1])) {
            return true
        }
    }
    fmt.Println("top right")
    // top right diagonal 
    if endIndex < len(arrayOfLines[lineIndex]) - 1 && lineIndex > 0 {
        if isSymbol(rune(arrayOfLines[lineIndex - 1][endIndex])) {
            return true
        }
    }
    fmt.Println("bottom left")
    // bottom left diagonal 
    if startIndex > 0 && lineIndex < len(arrayOfLines) - 1 {
        if isSymbol(rune(arrayOfLines[lineIndex + 1][startIndex - 1])) {
            return true
        }
    }


    // Up, down, diagonal cases
    for i := startIndex; i < endIndex; i++ {
        if lineIndex > 0 {
            if isSymbol(rune(arrayOfLines[lineIndex - 1][i])) {
                return true
            }
        }
        if lineIndex < len(arrayOfLines[lineIndex]) - 1 {
            if isSymbol(rune(arrayOfLines[lineIndex + 1][i])) {
                return true
            }
        }
    }

    return false
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    var arrayOfLines []string

    for scanner.Scan() {
        line := scanner.Text()
        arrayOfLines = append(arrayOfLines, line)
    }

    sum := 0 
    for lineIndex, line := range arrayOfLines  {
        for charIndex := 0; charIndex < len(line); charIndex++ {
            numDigitsInInt := 0

            startIndex := charIndex
            for charIndex < len(line) && unicode.IsDigit(rune(line[charIndex]))  {
                numDigitsInInt++
                charIndex++
            }
            
            // Found string of digits
            if(numDigitsInInt != 0) {
                endIndex := startIndex + numDigitsInInt
                fmt.Println(line[startIndex:endIndex])

                if hasAdjacentSymbol(lineIndex, charIndex, startIndex, endIndex, arrayOfLines) {
                    fmt.Println(line[startIndex:endIndex])
                    num, _ := strconv.Atoi(line[startIndex:endIndex])
                    sum += num
                }
            }

        }
        fmt.Println()

    }
    fmt.Println(sum)


}
