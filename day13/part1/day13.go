package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func patternLinesRowWise(pattern [][]byte) int {
	for i := 0; i < len(pattern)-1; i++ {
		if string(pattern[i]) == string(pattern[i+1]) {
			// Check if every line is equal
			if checkEquality(pattern, i, i+1) {
				return i + 1
			} else {
			}
		}
	}
	return 0
}

func checkEquality(pattern [][]byte, i, j int) bool {
	for shift := 1; shift <= i && i-shift >= 0 && j+shift < len(pattern); shift++ {
		if string(pattern[i-shift]) != string(pattern[j+shift]) {
			return false
		}
	}
	return true
}

func tranpose(slice [][]byte) [][]byte {
	rowLength := len(slice)
	colLength := len(slice[0])

	result := make([][]byte, colLength)
	for i := range result {
		result[i] = make([]byte, rowLength)
	}

	for i := 0; i < colLength; i++ {
		for j := 0; j < rowLength; j++ {
			result[i][j] = slice[j][i]
		}
	}

	return result
}
func main() {
	var patterns [][][]byte

	scanner := bufio.NewScanner(os.Stdin)
	var curPattern [][]byte
	for scanner.Scan() {
		line := scanner.Text()
		// seperate by newline
		if strings.TrimSpace(line) == "" {
			patterns = append(patterns, curPattern)
			curPattern = nil
			continue
		}
		curPattern = append(curPattern, []byte(line))
	}

	sum := 0
	for _, pattern := range patterns {
		// Transposing our matrix lets us do column-wise with our existing row-wise logic
		transposedPattern := tranpose(pattern)
		sum += patternLinesRowWise(transposedPattern) + 100*patternLinesRowWise(pattern)
	}
	fmt.Println(sum)
}
