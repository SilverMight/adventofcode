package main

import (
	"bufio"
	"fmt"
	"os"
)

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

func printPuzzle(slice [][]byte) {
	for _, v := range slice {
		fmt.Println(string(v))
	}
}

func roundAbovesCubes(puzzle string) map[int]int {
	roundsPerCube := make(map[int]int)

	currentCubeIndex := len(puzzle) + 1
	for i := 0; i < len(puzzle); i++ {
		switch puzzle[i] {
		case '#':
			currentCubeIndex = len(puzzle) - i

		case 'O':
			roundsPerCube[currentCubeIndex] += 1
		}
	}

	return roundsPerCube
}

func getSumPerCol(roundAboveCubes map[int]int) int {
	sum := 0
	for cube, rocks := range roundAboveCubes {
		for i := cube - 1; i >= cube-rocks; i-- {
			sum += i
		}
	}
	return sum
}

func getTotal(puzzle [][]byte) int {
	sum := 0
	for _, col := range puzzle {
		sum += getSumPerCol(roundAbovesCubes(string(col)))
	}

	return sum
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var puzzle [][]byte
	for scanner.Scan() {
		line := scanner.Text()
		puzzle = append(puzzle, []byte(line))
	}

	transposed := tranpose(puzzle)
	fmt.Println(getTotal(transposed))
}
