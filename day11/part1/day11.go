package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Vertex struct {
	X, Y int
}

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

func insertGalaxiesColumnWise(matrix [][]byte) [][]byte {
	for y := 0; y < len(matrix[0]); y++ {
		if !hasGalaxiesColumnwise(matrix, y) {
			fmt.Println("Column", y, "has no galaxies.")

			// insert .
			for x := 0; x < len(matrix); x++ {
				matrix[x] = append(matrix[x][:y], append([]byte{'.'}, matrix[x][y:]...)...)
			}
			// and skip this column, obviously
			y++

		}
	}

	return matrix
}

func hasGalaxiesColumnwise(matrix [][]byte, col int) bool {
	for x := 0; x < len(matrix); x++ {
		if matrix[x][col] == '#' {
			return true
		}
	}

	return false
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

func getShortestDistance(superpair []Vertex) int {
	pair1 := superpair[0]
	pair2 := superpair[1]

	deltaY := pair2.Y - pair1.Y
	deltaX := pair2.X - pair1.X

	distance := math.Abs(float64(deltaY)) + math.Abs(float64(deltaX))
	fmt.Println("From", pair1, "to", pair2, distance)
	return int(distance)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var matrix [][]byte
	for scanner.Scan() {
		line := []byte(scanner.Text())

		// No galaxies - add more ....
		if !strings.Contains(string(line), "#") {
			matrix = append(matrix, line)
		}

		matrix = append(matrix, line)
	}
	matrix = insertGalaxiesColumnWise(matrix)
	for _, row := range matrix {
		fmt.Println(string(row))
	}

	galaxies := getGalaxies(matrix)
	superpairs := getAllSuperPairs(galaxies)

	sum := 0
	for _, superpair := range superpairs {
		sum += getShortestDistance(superpair)
	}
	fmt.Println(sum)

}
