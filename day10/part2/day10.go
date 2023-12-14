package main

import (
	"bufio"
	"fmt"
	"os"
)

var pipeMap map[byte][]Vertex = map[byte][]Vertex{
	'|': {{1, 0}, {-1, 0}},
	'-': {{0, 1}, {0, -1}},
	'L': {{-1, 0}, {0, 1}},
	'J': {{-1, 0}, {0, -1}},
	'7': {{1, 0}, {0, -1}},
	'F': {{1, 0}, {0, 1}},
	'S': {{0, 1}, {0, -1}, {1, 0}, {-1, 0}},
}

type Vertex struct {
	X, Y int
}

func getStart(matrix [][]byte) Vertex {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 'S' {
				return Vertex{i, j}
			}
		}
	}
	return Vertex{-1, -1}
}

// BFS will get all the VALID tiles, i.e ignoring ..
func bfs(start Vertex, matrix [][]byte) [][]bool {
	rows, cols := len(matrix), len(matrix[0])
	queue := []Vertex{}

	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	sum := 0

	queue = append(queue, start)
	visited[start.X][start.Y] = true

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		sum++

		for _, direction := range pipeMap[matrix[curr.X][curr.Y]] {
			nextTile := Vertex{X: curr.X + direction.X, Y: curr.Y + direction.Y}
			if nextTile.X < 0 || nextTile.X >= rows || nextTile.Y < 0 || nextTile.Y >= cols {
				continue
			}
			if visited[nextTile.X][nextTile.Y] || matrix[nextTile.X][nextTile.Y] == '.' {
				continue
			}

			queue = append(queue, nextTile)
			visited[nextTile.X][nextTile.Y] = true

		}

	}

	return visited
}

func raycast(matrix [][]byte, visited [][]bool) (sum int) {
	// Hard coded change (very cringeworthy)
	start := getStart(matrix)

	matrix[start.X][start.Y] = 'F'

	for x, line := range matrix {
		enclosed := false
		var cornerchar rune
		for y, char := range line {
			// If this is actually part of the loop (what we care about)
			if visited[x][y] {
				switch char {
				case '.':
					if enclosed {
						sum++
						fmt.Print("I")
					} else {
						fmt.Print("O")
					}
				case '|':
					fmt.Print(string(char))
					enclosed = !enclosed
				// don't care
				case 'F':
					fmt.Print(string(char))
					cornerchar = 'F'
				case 'L':
					fmt.Print(string(char))
					cornerchar = 'L'
				case '7':
					fmt.Print(string(char))
					if cornerchar == 'L' {
						enclosed = !enclosed
					}
					cornerchar = ' '
				case 'J':
					fmt.Print(string(char))
					if cornerchar == 'F' {
						enclosed = !enclosed
					}
					cornerchar = ' '
				default:
					fmt.Print(string(char))
				}
			} else {
				if enclosed {
					fmt.Print("I")
					sum++
				} else {
					fmt.Print("O")
				}
			}
		}
		fmt.Println()
	}
	return
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var matrix [][]byte
	for scanner.Scan() {
		line := []byte(scanner.Text())
		matrix = append(matrix, line)
	}

	visited := bfs(getStart(matrix), matrix)
	fmt.Println(raycast(matrix, visited))

}
