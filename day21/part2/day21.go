package main

import (
	"bufio"
	"fmt"
	"os"
    "golang.org/x/exp/maps"
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

func floodfill(puzzle []string) map[Vertex]int {
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

    numSteps := len(puzzle)
    visited := make(map[Vertex]int)
    for i := 1; i <= numSteps ; i++ {

        // Handle only up to what's currently in our queue
        for remaining := len(queue); remaining > 0; remaining-- {
            curr := queue[0]
            queue = queue[1:]


            for _, dir := range directions {
                next := curr.add(dir)

                // Need to check presence in map before we do modulo on it
                if _, ok := visited[next]; ok {
                    if ok {
                        continue
                    }
                }

                if isValid(next) {
                    queue = append(queue, next)
                    visited[next] = i
                }
            }
        }
        fmt.Println(len(queue))
    }

    return visited
}

func part2(puzzle []string) int {
    visited := floodfill(puzzle)
    visitedValues := maps.Values(visited)

    // very taken from https://github.com/villuna/aoc23/wiki/A-Geometric-solution-to-advent-of-code-2023,-day-21
    var evenCorners int 
    var oddCorners int 

    var evenFull int
    var oddFull int

    for _, v := range visitedValues {
        if v % 2 == 0 {
            evenFull++
            if v > 65 {
                evenCorners++
            }
        }
        if v % 2 == 1 {
            oddFull++
            if v > 65 {
                oddCorners++
            }
        }
    }

    n := 202300

    total := ((n + 1) * (n + 1)) * oddFull + (n * n) * evenFull - (n+1) * oddCorners + n * evenCorners
    return  total
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

    fmt.Println(part2(puzzle))

}
