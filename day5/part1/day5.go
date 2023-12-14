package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type sourceToDest struct {
	source int64
	dest   int64
	length int64
}

func findMap(seed int64, sourceMap []sourceToDest) (loc int64) {
	for _, pair := range sourceMap {
		if seed >= pair.source && seed <= pair.source+(pair.length) {
			//fmt.Printf("Seed: %d, Source: %d, Dest %d, Length %d\n", seed, pair.source, pair.dest, pair.length)
			return (seed - pair.source) + pair.dest
		}

	}
	// no compatible pair
	return seed
}

func getLocation(seed int64, sourceMaps [][]sourceToDest) (loc int64) {
	currentLoc := seed
	for _, sourceMap := range sourceMaps {
		currentLoc = findMap(currentLoc, sourceMap)
	}

	return currentLoc
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	seeds := scanner.Text()
	seeds = seeds[strings.Index(seeds, ":")+1:]
	seedsArray := strings.Fields(seeds)
	fmt.Println(seedsArray)

	sourceDestMap := [][]sourceToDest{}
	previousLineContainsColon := false
	for scanner.Scan() {
		line := scanner.Text()
		previousLineContainsColon = strings.Contains(line, ":")

		// keep going while strings.Fields returns size  3
		if previousLineContainsColon {
			curMap := []sourceToDest{}
			for scanner.Scan() {
				nums := strings.Fields(scanner.Text())
				if len(nums) != 3 {
					break
				}

				source, _ := strconv.ParseInt(nums[1], 10, 64)
				dest, _ := strconv.ParseInt(nums[0], 10, 64)
				length, _ := strconv.ParseInt(nums[2], 10, 64)

				curMap = append(curMap, sourceToDest{source: source, dest: dest, length: int64(length)})
			}

			sourceDestMap = append(sourceDestMap, curMap)
		}

	}

	var minimumLocation int64 = math.MaxInt64
	for _, seed := range seedsArray {
		seed64, _ := strconv.ParseInt(seed, 10, 64)
		curLoc := getLocation(seed64, sourceDestMap)
		if curLoc < minimumLocation {
			minimumLocation = curLoc
		}
	}
	fmt.Println(minimumLocation)
}
