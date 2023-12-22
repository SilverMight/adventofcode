package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
    Rock = '#'
    Plot = '.'
    Start = 'S'
)

type Vertex struct {
    X, Y int
}

func (v1 Vertex) add(v2 Vertex) Vertex {
    return Vertex{X: v1.X + v2.X, Y: v1.Y + v2.Y}
}

var directions []Vertex = []Vertex{
    {1, 0},
    {-1, 0},
    {0, 1},
    {0, -1},
}

func findStart(puzzle []string) Vertex {
    for x := range puzzle {
        for y := range puzzle {
            if puzzle[x][y] == Start {
                return Vertex{x, y}
            }
        }
    }
    return Vertex{-1, -1}
}


func part1(puzzle []string) int {
    queue := []Vertex{findStart(puzzle)}

    isValid := func (v Vertex) bool {
        if v.X < 0 || v.X >= len(puzzle) || v.Y < 0 || v.Y >= len(puzzle[0]) {
            return false
        }
        if puzzle[v.X][v.Y] == Rock {
            return false
        }
        return true
    }

    fillVisitedArray := func () [][]bool {
        visited := make([][]bool, len(puzzle))
        for i := range visited {
            visited[i] = make([]bool, len(puzzle[0])) 
        }
        return visited
    }

    numSteps := 64
    for i := 0; i < numSteps ; i++ {

        // We only care about revisiting the same tile in a single step
        // We could revisit the tile after.
        visited := fillVisitedArray() 

        // Handle only up to what's currently in our queue
        for remaining := len(queue); remaining > 0; remaining-- {
            curr := queue[0]
            queue = queue[1:]


            for _, dir := range directions {
                next := curr.add(dir)

                if isValid(next) && !visited[next.X][next.Y] {
                    queue = append(queue, next)
                    visited[next.X][next.Y] = true
                }
            }
        }
    }

    return len(queue)
}

func main() {
    input, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Failed to open file", err)
    }
    scanner := bufio.NewScanner(input)
    var puzzle []string
    for scanner.Scan() {
        line := scanner.Text()
        puzzle = append(puzzle, line)
    } 

    fmt.Println(part1(puzzle))
}
