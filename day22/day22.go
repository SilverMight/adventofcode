package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Vertex3D struct {
    X, Y, Z int
}

type Brick struct {
    start Vertex3D
    end Vertex3D
}

// Slice must be sorted by Z
func brickColliding(target Brick, bricksUnder []Brick) (newBrick Brick, supportingBricks []Brick) {
    testFallBrick := target
    for testFallBrick.start.Z > 1 {
        collisionFound := false
        for _, brick := range bricksUnder {
            // possible collision (on same level)
            if (testFallBrick.start.Z - 1 <= brick.end.Z) && (testFallBrick.end.Z - 1 >= brick.start.Z) {
                // they overlap
                if (testFallBrick.start.X <= brick.end.X) && (testFallBrick.end.X >= brick.start.X) && (testFallBrick.start.Y <= brick.end.Y) && (testFallBrick.end.Y >= brick.start.Y) {
                    supportingBricks = append(supportingBricks, brick)
                    collisionFound = true
                }
            }
            
        }
        // If we already found a collision, we can't fall anymore
        if collisionFound {
            return testFallBrick, supportingBricks
        }
        testFallBrick.end.Z--
        testFallBrick.start.Z--
    }
    
    // we just fell all the way, no collisions
    return testFallBrick, supportingBricks
}

func part1(bricks []Brick, bricksToSupportingBricks map[Brick][]Brick) int {
    
    disintegratable := len(bricks)

    // I really just want a set here, but go doesn't have that lol
    isExclusivelySupportingBricks := make(map[Brick]bool)

    for _, brick := range bricks {
        // This brick is exclusively supported by this brick
        if len(bricksToSupportingBricks[brick]) == 1  {
            isExclusivelySupportingBricks[bricksToSupportingBricks[brick][0]] = true
        }
    }

    for _, ok := range isExclusivelySupportingBricks {
        if ok {
            disintegratable--
        } 
    }

    return disintegratable
}

func part2(bricks []Brick, bricksBelow map[Brick][]Brick) int {
    bricksAbove := make(map[Brick][]Brick)
    for brick, bricksBelow := range bricksBelow {
        for _, below := range bricksBelow {
            bricksAbove[below] = append(bricksAbove[below], brick) 
        }
    }

    sum := 0 

    allBricksFallen := func(fallen map[Brick]bool, bricksBelowUs []Brick) bool {
        for _, brick := range bricksBelowUs {
            if _, hasFallen := fallen[brick]; !hasFallen {
                return false 
            }
        }
        return true
    }

    for _, brick := range bricks {
        fallen := make(map[Brick]bool)
        fallen[brick] = true
        queue := []Brick{brick}

        for len(queue) > 0 {
            curr := queue[0]
            queue = queue[1:]

            // Check everything above us (what we're supporting)
            for _, supporting := range bricksAbove[curr] {
                // if everything below us fell
                if allBricksFallen(fallen, bricksBelow[supporting]) {
                    fallen[supporting] = true
                    queue = append(queue, supporting)
                }
            }
        }

        sum += len(fallen) - 1
    }
    return sum
}

func brickFalling(bricks []Brick) ([]Brick, map[Brick][]Brick) {
    brickToSupportingBricks := make(map[Brick][]Brick)
    for i, brick := range bricks {
        var supportingBricks []Brick
        bricks[i], supportingBricks = brickColliding(brick, bricks[:i]) 
        brickToSupportingBricks[bricks[i]] = supportingBricks
    }
    
    return bricks, brickToSupportingBricks
}

func parse(puzzle []string) []Brick {
    var bricks []Brick
    for _, line := range puzzle {
        startAndEnd := strings.Split(line, "~")

        startArray := strings.Split(startAndEnd[0], ",")
        endArray := strings.Split(startAndEnd[1], ",")

        startX, _ := strconv.Atoi(startArray[0])
        startY, _ := strconv.Atoi(startArray[1])
        startZ, _ := strconv.Atoi(startArray[2])

        endX, _ := strconv.Atoi(endArray[0])
        endY, _ := strconv.Atoi(endArray[1])
        endZ, _ := strconv.Atoi(endArray[2])
        startCoords := Vertex3D{X: startX, Y: startY, Z: startZ}
        endCoords := Vertex3D{X: endX, Y: endY, Z: endZ}

        bricks = append(bricks, Brick{start: startCoords, end: endCoords})

    }

    return bricks
}

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println("Failed to open file:", err)

    }
    scanner := bufio.NewScanner(file)
    
    var puzzle []string
    for scanner.Scan() {
        line := scanner.Text()

        puzzle = append(puzzle, line)
    }
    bricks := parse(puzzle)

    // Sort by Z
    slices.SortFunc(bricks, func (a, b Brick) int {
        return min(a.start.Z, a.end.Z) - min(b.start.Z, b.end.Z)
    })

    fallenBricks, bricksBelow := brickFalling(bricks)
    fmt.Println("Part 1", part1(fallenBricks, bricksBelow))
    fmt.Println("Part 2", part2(fallenBricks, bricksBelow))
}
