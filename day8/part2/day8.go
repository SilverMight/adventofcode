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

func numStepsToEnd(currentNode string, nodeMap map[string][]string, instructions string) int {
	directionNum := LEFT
	numSteps := 0
	for true {
		for _, direction := range instructions {

			if direction == 'L' {
				directionNum = LEFT
			}
			if direction == 'R' {
				directionNum = RIGHT
			}
			//fmt.Println("Moving", currentNode, nodeMap[currentNode][directionNum])
			currentNode = nodeMap[currentNode][directionNum]

			numSteps++

			if currentNode[2] == 'Z' {
				return numSteps
			}
		}
	}

	return 0
}

func traverseNodeMap(nodeMap map[string][]string, instructions string, startingNodes []string) []int {
	currentNodes := startingNodes
	nodeCycles := make([]int, len(startingNodes))
	for nodeIndex, startingNode := range currentNodes {
		fmt.Println(nodeIndex)
		nodeCycles[nodeIndex] = numStepsToEnd(startingNode, nodeMap, instructions)
	}

	return nodeCycles
}

func findLCM(numbers []int) int {
	lcm := numbers[0]

	// Iterate through the array and find the LCM of each element with the current LCM
	for i := 1; i < len(numbers); i++ {
		lcm = calculateLCM(lcm, numbers[i])
	}

	return lcm
}

// calculateLCM calculates the LCM of two integers using the GCD (Greatest Common Divisor) formula
func calculateLCM(a, b int) int {
	return a * b / calculateGCD(a, b)
}

// calculateGCD calculates the Greatest Common Divisor (GCD) of two integers using Euclidean algorithm
func calculateGCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
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

	cycles := traverseNodeMap(nodeMap, instructions, startingNodes)
	fmt.Println(findLCM(cycles))

}
