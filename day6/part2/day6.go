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

	var maxTime, recordDistance int

	for i := 0; i < 2; i++ {
		scanner.Scan()
		line := scanner.Text()
		line = line[strings.Index(line, ":")+1:]
		input := strings.Join(strings.Fields(line), "")

		if i == 0 {
			maxTime, _ = strconv.Atoi(input)
		} else {
			recordDistance, _ = strconv.Atoi(input)
		}
	}

	waysToWin := 0
	for time := 1; time < maxTime; time++ {
		if canWin(maxTime, recordDistance, time) {
			waysToWin++
		}
	}

	fmt.Println("\n", waysToWin)
}
