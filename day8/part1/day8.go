package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	LEFT  = 0
	RIGHT = 1
)

func traverseNodeMap(nodeMap map[string][]string, instructions string) int {
	numSteps := 0
	currentNode := "AAA"
	for true {
		for _, direction := range instructions {
			if direction == 'L' {
				currentNode = nodeMap[currentNode][LEFT]
			}
			if direction == 'R' {
				currentNode = nodeMap[currentNode][RIGHT]
			}

			numSteps++

			if currentNode == "ZZZ" {
				return numSteps
			}
		}

	}

	return numSteps
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	instructions := scanner.Text()

	scanner.Scan()

	nodeMap := make(map[string][]string)
	for scanner.Scan() {
		line := scanner.Text()

		node := line[0:3]

		leftElement := line[7:10]
		rightElement := line[12:15]

		nodeMap[node] = append(nodeMap[node], leftElement)
		nodeMap[node] = append(nodeMap[node], rightElement)
	}

	fmt.Println(traverseNodeMap(nodeMap, instructions))

}
