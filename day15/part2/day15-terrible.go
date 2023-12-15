// WOULD work, if go maps preserved insertion order, but they don't.
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

type lens struct {
	focus int
	index int
}

func getPower(boxes []map[string]int) (power int) {
	for boxNumber, box := range boxes {
		if len(box) == 0 {
			continue
		}

		index := 1
		fmt.Println(boxNumber, box)
		for _, lens := range box {
			fmt.Println(boxNumber+1, "*", index, "*", lens)
			power += (boxNumber + 1) * (index) * lens
			index++
		}

	}
	return power
}

func calculateSteps(boxes []map[string]int, steps []string) {
	for _, step := range steps {
		// =
		if index := strings.Index(step, "="); index != -1 {
			label := step[:index]
			amount := int(step[index+1] - '0')

			boxes[hash(label)][label] = amount
		}
		if index := strings.Index(step, "-"); index != -1 {
			label := step[:index]
			if _, ok := boxes[hash(label)][label]; ok {
				delete(boxes[hash(label)], label)
			}
		}

		for boxNumber, box := range boxes {
			if len(box) == 0 {
				continue
			}
			fmt.Println(boxNumber, box)
		}
	}
	fmt.Println(getPower(boxes))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var steps []string
	boxes := make([]map[string]int, 256)
	for i := range boxes {
		boxes[i] = make(map[string]int)
	}
	for scanner.Scan() {
		line := scanner.Text()
		steps = strings.Split(line, ",")
	}
	calculateSteps(boxes, steps)

}
