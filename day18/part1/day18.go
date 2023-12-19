package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Direction byte 

const (
    Up Direction = 'U'
    Down Direction = 'D'
    Left Direction = 'L'
    Right Direction = 'R'
)

type instruction struct {
    dir Direction
    moves int
    RGB string
}

type Vertex struct {
    x, y int
}

// Shoelace theorem
func getInsideArea(corners []Vertex) (int) {
    area := 0
    length := len(corners)

    for i := 0; i < length - 1; i++ {
        area += corners[i].x * corners[(i+1) % length].y - corners[(i + 1) % length].x * corners[i].y 
    }
    if (area < 0) {
        return -area / 2
    }
    return area / 2
}

func getCornersAndPerimeter(instructions []instruction) ([]Vertex, int){
    lastX, lastY := 0, 0
    corners := []Vertex{{lastX, lastY}}

    perimeter := 0
    for _, instr := range instructions {
        switch instr.dir {
        case Up:
            lastX -= instr.moves 
        case Down:
            lastX += instr.moves
        case Left: 
            lastY -= instr.moves
        case Right: 
            lastY += instr.moves
        }

        corners = append(corners, Vertex{lastX, lastY})
        perimeter += instr.moves
    }

    return corners, perimeter
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    var instructions []instruction
    for scanner.Scan() {
        line := scanner.Text()
        fields := strings.Fields(line)
        direction := Direction(fields[0][0])
        moves, _ := strconv.Atoi(fields[1])
        instructions = append(instructions, instruction{direction, moves, fields[2]})
    }

    corners, perimeter := getCornersAndPerimeter(instructions)
    area := getInsideArea(corners)

    // Pick's theorem
    fmt.Println(area + (perimeter / 2) + 1)

}
