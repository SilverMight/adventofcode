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

func getStart(matrix []string) Vertex {
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
func bfs(start Vertex, matrix []string) (distance int) {
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

	distance = sum / 2
	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var matrix []string
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, line)
	}

	startVertex := getStart(matrix)
	fmt.Println(bfs(startVertex, matrix))

}
