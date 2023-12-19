package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Direction int 

const (
    Up Direction = 3
    Down Direction = 1
    Left Direction = 2
    Right Direction = 0
)

type instruction struct {
    dir Direction
    moves int
}

type Vertex struct {
    x, y int
}

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

// Shoelace theorem
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

func parseColorToInstructions(color string) instruction {
    fmt.Println(color)
    distanceStr := color[0:5]
    directionNum := int(color[5] - '0')
    distanceNum, _ := strconv.ParseUint(distanceStr, 16, 64)
    fmt.Println(directionNum, distanceNum)

    return instruction{Direction(directionNum), int(distanceNum)}
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    var instructions []instruction
    for scanner.Scan() {
        line := scanner.Text()
        fields := strings.Fields(line)
        color := fields[2][2:8]

        
        instr := parseColorToInstructions(color)
        instructions = append(instructions, instr)
    }

    corners, perimeter := getCornersAndPerimeter(instructions)
    area := getInsideArea(corners)

    // Pick's theorem
    fmt.Println(area + (perimeter / 2) + 1)

}
