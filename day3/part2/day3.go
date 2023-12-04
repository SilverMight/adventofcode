package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type indexPair struct {
    row int
    col int
}
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

func getAdjacentSymbolIndex(lineIndex, charIndex, startIndex, endIndex int, arrayOfLines []string) (hasAdjacent bool, pair indexPair) {


    // To the left of us
    if startIndex > 0 {
        if isSymbol(rune(arrayOfLines[lineIndex][startIndex - 1])) {
            return true, indexPair{lineIndex, startIndex - 1}
        }
    }

    // To the right of us
    if endIndex < len(arrayOfLines[lineIndex]) {
        if isSymbol(rune(arrayOfLines[lineIndex][endIndex])) {
            return true, indexPair{lineIndex, endIndex}
        }
    }

    // bottom right diagonal
    if endIndex < len(arrayOfLines[lineIndex]) - 1 && lineIndex < len(arrayOfLines) - 1 {
        if isSymbol(rune(arrayOfLines[lineIndex + 1][endIndex])) {
            return true, indexPair{lineIndex + 1, endIndex}
        }
    }

    // top left diagonal 
    if startIndex > 0 && lineIndex > 0 {
        if isSymbol(rune(arrayOfLines[lineIndex - 1][startIndex - 1])) {
            return true, indexPair{lineIndex - 1, startIndex - 1}
        }
    }

    // top right diagonal 
    if endIndex < len(arrayOfLines[lineIndex]) - 1 && lineIndex > 0 {
        if isSymbol(rune(arrayOfLines[lineIndex - 1][endIndex])) {
            return true, indexPair{lineIndex - 1, endIndex}
        }
    }

    // bottom left diagonal 
    if startIndex > 0 && lineIndex < len(arrayOfLines) - 1 {
        if isSymbol(rune(arrayOfLines[lineIndex + 1][startIndex - 1])) {
            return true, indexPair{lineIndex + 1, startIndex - 1}
        }
    }


    // Up, down, diagonal cases
    for i := startIndex; i < endIndex; i++ {
        if lineIndex > 0 {
            if isSymbol(rune(arrayOfLines[lineIndex - 1][i])) {
                return true, indexPair{lineIndex - 1, i}
            }
        }
        if lineIndex < len(arrayOfLines[lineIndex]) - 1 {
            if isSymbol(rune(arrayOfLines[lineIndex + 1][i])) {
                return true, indexPair{lineIndex + 1, i}
            }
        }
    }

    return false, indexPair{-1, -1}
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    indexPairToValue := make(map[indexPair][]int)

    
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

                if hasAdjacent, pair := getAdjacentSymbolIndex(lineIndex, charIndex, startIndex, endIndex, arrayOfLines); hasAdjacent {
                    fmt.Println(line[startIndex:endIndex])
                    value, _ := strconv.Atoi(line[startIndex:endIndex])
                    indexPairToValue[pair] = append(indexPairToValue[pair], value)
                }
            }

        }

    }
    for pair := range indexPairToValue {
        // If we have exactly two adjacent symbols
        if(len(indexPairToValue[pair]) == 2) {
            sum += indexPairToValue[pair][0] * indexPairToValue[pair][1]
        }
    }

    fmt.Println(sum)
}
