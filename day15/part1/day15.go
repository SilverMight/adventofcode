package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func hash(step string) (val int) {
	for _, c := range step {
		val += int(c)
		val *= 17
		val %= 256
	}

	return val
}

func sumHashedSteps(steps []string) (sum int) {
	for _, step := range steps {
		sum += hash(step)
		fmt.Println(sum)
	}

	return sum
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var steps []string
	for scanner.Scan() {
		line := scanner.Text()
		steps = strings.Split(line, ",")
	}

	fmt.Println(sumHashedSteps(steps))

}
