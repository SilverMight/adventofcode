package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func patternLinesRowWise(pattern [][]byte) int {
	printMatrix(pattern)
	for i := 0; i < len(pattern)-1; i++ {
		// If there is a smudge on our line of reflection
		if differingCharacters(pattern[i], pattern[i+1]) == 1 {
			// equality check CAN not have smudge
			if checkEquality(pattern, i, i+1, 0) {
				return i + 1
			}
		}
		// If there is no smudge on our line of reflection
		if differingCharacters(pattern[i], pattern[i+1]) == 0 {
			// equality check CAN have smudge
			if checkEquality(pattern, i, i+1, 1) {
				return i + 1
			}
		}
	}
	return 0
}

func differingCharacters(str1, str2 []byte) (differingChars int) {
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			differingChars++
		}
	}
	return
}

func fixSmudge(str1, str2 []byte) (differingchars int) {
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			str1[i] = str2[i]
		}
	}
	return
}

func checkEquality(pattern [][]byte, i, j int, smudge int) bool {
	differingChars := 0
	for shift := 1; shift <= i && i-shift >= 0 && j+shift < len(pattern); shift++ {
		// Does not Equals
		differingChars += differingCharacters(pattern[i-shift], pattern[j+shift])
	}
	// 1 character if we are allowed to smudge, 0 if not
	return differingChars == smudge
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

func printMatrix(matrix [][]byte) {
	for _, line := range matrix {
		fmt.Println(string(line))
	}

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
