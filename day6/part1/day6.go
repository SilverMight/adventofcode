package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func canWin(maxTime, minimumDistance, buttonHoldTime int) bool {
	speed := buttonHoldTime
	distanceTraveled := speed * (maxTime - buttonHoldTime)
	if distanceTraveled > minimumDistance {
		return true
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var times, distances []string

	for i := 0; i < 2; i++ {
		scanner.Scan()
		line := scanner.Text()
		line = line[strings.Index(line, ":")+1:]
		fields := strings.Fields(line)

		if i == 0 {
			times = fields
		} else {
			distances = fields
		}
	}

	power := 1

	for i := 0; i < len(times); i++ {
		currentDistance, _ := strconv.Atoi(distances[i])
		currentMaxTime, _ := strconv.Atoi(times[i])

		waysToWin := 0
		for time := 1; time < currentMaxTime; time++ {
			if canWin(currentMaxTime, currentDistance, time) {
				waysToWin++
			}
		}

		power *= waysToWin
	}
	fmt.Println("\n", power)
}
