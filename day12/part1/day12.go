package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	OPERATIONAL = '.'
	DAMAGED     = '#'
	UNKNOWN     = '?'
)

func getPossibleArrangements(springs string, damagedGroups []int, curDamaged int) int {

	// hit all our groups
	if len(springs) == 0 {
		if len(damagedGroups) == 0 && curDamaged == 0 {
			return 1
		} else if len(damagedGroups) == 1 && damagedGroups[0] == curDamaged {
			return 1
		} else {
			return 0
		}
	}

	if springs[0] == UNKNOWN {
		return getPossibleArrangements(strings.Replace(springs, "?", ".", 1), damagedGroups, curDamaged) + getPossibleArrangements(strings.Replace(springs, "?", "#", 1), damagedGroups, curDamaged)
	}

	if springs[0] == DAMAGED {
		return getPossibleArrangements(springs[1:], damagedGroups, curDamaged+1)

	}

	// Just keep moving
	if springs[0] == OPERATIONAL {
		// Correct size
		if len(damagedGroups) > 0 && curDamaged == damagedGroups[0] {
			return getPossibleArrangements(springs[1:], damagedGroups[1:], 0)
			// No streak
		} else if curDamaged == 0 {
			return getPossibleArrangements(springs[1:], damagedGroups, 0)
		} else {
			return 0
		}
	}

	return 0
}

func stringsToIntegers(array []string) []int {
	ret := make([]int, len(array))
	for index, number := range array {
		ret[index], _ = strconv.Atoi(number)
	}

	return ret
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		numOfDamagedSprings := stringsToIntegers(strings.Split(fields[1], ","))

		sum += getPossibleArrangements(fields[0], numOfDamagedSprings, 0)
	}

	fmt.Println(sum)
}
