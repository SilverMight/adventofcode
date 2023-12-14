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

func allNodesEnding(currentNodes []string) bool {
	for _, node := range currentNodes {
		if node[2] != 'Z' {
			return false
		}
	}

	return true
}

func moveNodes(currentNodes []string, nodeMap map[string][]string, direction rune) []string {
	directionNum := LEFT
	if direction == 'L' {
		directionNum = LEFT
	}
	if direction == 'R' {
		directionNum = RIGHT
	}

	for nodeNum, node := range currentNodes {
		fmt.Println("Moving", node, "to", nodeMap[node][directionNum])
		currentNodes[nodeNum] = nodeMap[node][directionNum]
	}

	return currentNodes
}

func traverseNodeMap(nodeMap map[string][]string, instructions string, startingNodes []string) int {
	numSteps := 0
	currentNodes := startingNodes
	for true {
		for _, direction := range instructions {

			currentNodes = moveNodes(currentNodes, nodeMap, direction)

			numSteps++

			if allNodesEnding(currentNodes) {
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
	var startingNodes []string
	for scanner.Scan() {
		line := scanner.Text()

		node := line[0:3]

		// starting node
		if line[2] == 'A' {
			startingNodes = append(startingNodes, node)
		}

		leftElement := line[7:10]
		rightElement := line[12:15]

		nodeMap[node] = append(nodeMap[node], leftElement)
		nodeMap[node] = append(nodeMap[node], rightElement)
	}

	fmt.Println(traverseNodeMap(nodeMap, instructions, startingNodes))

}
