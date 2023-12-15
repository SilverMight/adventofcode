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
const (
	NUMCOPIES = 5
)

type memoize struct {
	springs           string
	damagedGroupsLeft string
	curDamaged        int
}

var solved map[memoize]int = make(map[memoize]int)

// I know this is ridiculous. please don't tell me
func damagedGroupsString(array []int) string {
	groups := make([]string, len(array))
	for i, v := range array {
		groups[i] = strconv.Itoa(v)
	}
	return strings.Join(groups, ",")
}

func getPossibleArrangements(springs string, damagedGroups []int, curDamaged int) int {
	key := memoize{springs, damagedGroupsString(damagedGroups), curDamaged}

	// If we solved it, use it
	if val, inMap := solved[key]; inMap {
		return val
	}

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

	// ? - either, try # or .
	if springs[0] == UNKNOWN {
		result := getPossibleArrangements(strings.Replace(springs, "?", ".", 1), damagedGroups, curDamaged) + getPossibleArrangements(strings.Replace(springs, "?", "#", 1), damagedGroups, curDamaged)
		// Save our result
		solved[key] = result
		return result
	}

	// Move up and count another damaged spring
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
		var newSprings string
		// copies
		for i := 0; i < 5; i++ {
			newSprings += fields[0]
			if i == 4 {
				break
			}
			newSprings += "?"
		}

		var newNumOfDamagedSprings []int
		for i := 0; i < 5; i++ {
			newNumOfDamagedSprings = append(newNumOfDamagedSprings, numOfDamagedSprings...)
		}

		sum += getPossibleArrangements(newSprings, newNumOfDamagedSprings, 0)

	}
	fmt.Println(sum)
}
