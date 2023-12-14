package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Vertex struct {
	X, Y int
}

const (
	EXPANSION = 1000000
)

func getGalaxies(matrix [][]byte) (galaxies []Vertex) {
	for x := 0; x < len(matrix); x++ {
		for y := 0; y < len(matrix[0]); y++ {
			if matrix[x][y] == '#' {
				galaxies = append(galaxies, Vertex{x, y})
			}
		}
	}

	return
}

func hasGalaxiesColumnwise(matrix [][]byte, col int) bool {
	for x := 0; x < len(matrix); x++ {
		if matrix[x][col] == '#' {
			return true
		}
	}

	return false
}

func hasGalaxiesRowwise(matrix [][]byte, row int) bool {
	return strings.Contains(string(matrix[row]), "#")
}

func getAllSuperPairs(galaxies []Vertex) (superpairs [][]Vertex) {
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			superPair := []Vertex{galaxies[i], galaxies[j]}
			superpairs = append(superpairs, superPair)
		}
	}

	return
}

func getShortestDistance(matrix [][]byte, superpair []Vertex) int {
	pair1 := superpair[0]
	pair2 := superpair[1]

	distance := rowDistance(matrix, pair1.X, pair2.X) + colDistance(matrix, pair1.Y, pair2.Y)
	return distance
}

func rowDistance(matrix [][]byte, x, y int) (distance int) {
	var start, end int
	if x < y {
		start = x
		end = y
	} else {
		start = y
		end = x
	}
	for row := start + 1; row <= end; row++ {
		if hasGalaxiesRowwise(matrix, row) {
			distance += 1
		} else {
			distance += EXPANSION
		}
	}
	return distance
}

func colDistance(matrix [][]byte, x, y int) (distance int) {
	var start, end int
	if x < y {
		start = x
		end = y
	} else {
		start = y
		end = x
	}
	for col := start + 1; col <= end; col++ {
		if hasGalaxiesColumnwise(matrix, col) {
			distance += 1
		} else {
			distance += EXPANSION
		}
	}

	return distance
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var matrix [][]byte
	for scanner.Scan() {
		line := []byte(scanner.Text())

		matrix = append(matrix, line)
	}

	galaxies := getGalaxies(matrix)
	superpairs := getAllSuperPairs(galaxies)

	sum := 0
	for _, superpair := range superpairs {
		sum += getShortestDistance(matrix, superpair)
	}
	fmt.Println(sum)

}
