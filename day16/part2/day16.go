package main

import (
	"bufio"
	"fmt"
	"os"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Vertex struct {
	X, Y int
	dir  Direction
}

func getNext(v1 Vertex) Vertex {
	switch v1.dir {
	case Up:
		return Vertex{X: v1.X - 1, Y: v1.Y, dir: Up}
	case Down:
		return Vertex{X: v1.X + 1, Y: v1.Y, dir: Down}
	case Left:
		return Vertex{X: v1.X, Y: v1.Y - 1, dir: Left}
	case Right:
		return Vertex{X: v1.X, Y: v1.Y + 1, dir: Right}
	}

	return Vertex{-1, -1, Right}
}

func splitter(splitter byte, v1 Vertex) []Vertex {
	switch splitter {
	case '-':
		switch v1.dir {
		case Left, Right:
			return []Vertex{getNext(v1)}
		case Up, Down:
			leftVertex, rightVertex := v1, v1
			leftVertex.dir = Left
			rightVertex.dir = Right
			return []Vertex{getNext(leftVertex), getNext(rightVertex)}

		}
	case '|':
		switch v1.dir {
		case Up, Down:
			return []Vertex{getNext(v1)}
		case Left, Right:
			upVertex, downVertex := v1, v1
			upVertex.dir = Up
			downVertex.dir = Down
			return []Vertex{getNext(upVertex), getNext(downVertex)}
		}
	}

	return []Vertex{}
}

func mirror(mirror byte, v1 Vertex) Vertex {
	switch mirror {
	case '/':
		switch v1.dir {
		case Right:
			v1.dir = Up
		case Left:
			v1.dir = Down
		case Up:
			v1.dir = Right
		case Down:
			v1.dir = Left
		}
	case '\\':
		switch v1.dir {
		case Right:
			v1.dir = Down
		case Left:
			v1.dir = Up
		case Up:
			v1.dir = Left
		case Down:
			v1.dir = Right
		}
	}
	return getNext(v1)
}

func traversal(puzzle []string, start Vertex) int {
	visited := make(map[Vertex]bool)
	visitedDirectionless := make(map[Vertex]bool)

	queue := []Vertex{}

	queue = append(queue, start)

	isValid := func(point Vertex, mirror bool) bool {
		if point.X < 0 || point.X >= len(puzzle) || point.Y < 0 || point.Y >= len(puzzle[0]) {
			return false
		}
		if visited[point] {
			return false
		}

		return true
	}

	sum := 0
	// Standard BFS
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		sum++

		curType := puzzle[curr.X][curr.Y]
		visited[curr] = true

		// Directionless emplace
		dirrectionlessCurr := curr
		dirrectionlessCurr.dir = Up // make everything up (not unique)
		visitedDirectionless[dirrectionlessCurr] = true

		switch curType {
		case '.':
			next := getNext(curr)
			next.dir = curr.dir
			if isValid(next, false) {
				queue = append(queue, next)
			}
		case '-', '|':
			nextArray := splitter(curType, curr)
			for _, next := range nextArray {
				if isValid(next, true) {
					queue = append(queue, next)
				}
			}
		case '/', '\\':
			next := mirror(curType, curr)
			if isValid(next, true) {
				queue = append(queue, next)
			}
		}
	}

	return len(visitedDirectionless)
}

func getBestSum(puzzle []string) int {
	rows, cols := len(puzzle), len(puzzle[0])
	cornerCases := []Vertex{
		// top left
		{0, 0, Right},
		{0, 0, Down},
		// top right
		{0, cols - 1, Left},
		{0, cols - 1, Down},
		// bottom left
		{rows - 1, 0, Right},
		{rows - 1, 0, Up},
		// bottom right
		{rows - 1, cols - 1, Left},
		{rows - 1, cols - 1, Up},
	}

	best := -1

	// All edge tiles
	updateBest := func(start Vertex) {
		currentSum := traversal(puzzle, start)
		if currentSum > best {
			best = currentSum
		}
	}

	for _, corner := range cornerCases {
		updateBest(corner)
	}

	// Iterate over edge tiles
	for i := 0; i < cols; i++ {
		updateBest(Vertex{0, i, Down})      // Top row
		updateBest(Vertex{rows - 1, i, Up}) // Bottom row
	}
	for i := 0; i < rows; i++ {
		updateBest(Vertex{i, 0, Right})       // Left column
		updateBest(Vertex{i, cols - 1, Left}) // Right column
	}
	return best
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var puzzle []string
	for scanner.Scan() {
		line := scanner.Text()
		puzzle = append(puzzle, line)
	}

	//sum := traversal(puzzle, Vertex{0, 0, Right})
	best := getBestSum(puzzle)
	fmt.Println(best)
}
