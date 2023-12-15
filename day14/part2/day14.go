// This is definitely not very efficient
package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func gridToString(puzzle [][]byte) []byte {
	var sol []byte
	for i := range puzzle {
		sol = append(sol, puzzle[i]...)
	}
	return sol
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

func printPuzzle(slice [][]byte) {
	for _, v := range slice {
		fmt.Println(string(v))
	}
}

func slide(puzzle [][]byte, shouldReverse bool) [][]byte {

	slid := make([][]byte, len(puzzle))

	if shouldReverse {
		reverse(puzzle)
	}

	for line := range puzzle {
		currentCubeIndex := -1
		var newLine []byte
		for i, block := range puzzle[line] {
			switch block {
			case '#':
				newLine = append(newLine, block)
				currentCubeIndex = i
			case '.':
				newLine = append(newLine, block)
			case 'O':
				newLine = slices.Insert(newLine, currentCubeIndex+1, block)
				currentCubeIndex++
			}
		}
		slid[line] = newLine
	}
	if shouldReverse {
		reverse(slid)
	}

	//printPuzzle(tranpose(slid))
	return slid

}

func reverse(puzzle [][]byte) {
	for i := range puzzle {
		slices.Reverse(puzzle[i])
	}
}

func tilt(puzzle [][]byte, dir Direction) [][]byte {
	transposed := tranpose(puzzle)
	switch dir {
	case North:
		return tranpose(slide(transposed, false))
	case East:
		slid := slide(puzzle, true)
		return slid
	case West:
		slid := slide(puzzle, false)
		return slid
	case South:
		slid := slide(transposed, true)
		return tranpose(slid)
	}

	return transposed
}

func spin(puzzle [][]byte) [][]byte {
	cycle := puzzle
	for _, dir := range []Direction{North, West, South, East} {
		cycle = tilt(cycle, dir)
	}
	return cycle
}

func count(puzzle [][]byte) int {
	sum := 0
	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle[0]); j++ {
			if puzzle[i][j] == 'O' {
				sum += len(puzzle) - i
			}
		}
	}
	return sum
}

func part2(puzzle [][]byte) int {
	spun := puzzle
	storedGrids := make(map[string]int)
	iters := 1000000000
	cycleStart := 0
	i := 0
	for ; i < iters; i++ {
		spun = spin(spun)
		printPuzzle(spun)
		stringifiedSpin := string(gridToString(spun))

		fmt.Println(count(spun))
		// Found cycle
		if v, ok := storedGrids[stringifiedSpin]; ok {
			cycleStart = v
			break
		}

		storedGrids[stringifiedSpin] = i

	}
	// See amount remaining in this cycle
	cyclesLeft := (iters - i - 1) % (i - cycleStart)

	for cycle := 0; cycle < cyclesLeft; cycle++ {
		spun = spin(spun)
	}
	return count(spun)

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var puzzle [][]byte
	for scanner.Scan() {
		line := scanner.Text()
		puzzle = append(puzzle, []byte(line))
	}

	fmt.Println(part2(puzzle))
}
