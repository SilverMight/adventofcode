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
	label       string
	focallength int
}

func getPower(boxes [][]lens) (power int) {
	for boxNumber, box := range boxes {
		if len(box) == 0 {
			continue
		}

		fmt.Println(boxNumber, box)
		for i, lens := range box {
			power += (boxNumber + 1) * (i + 1) * lens.focallength
		}

	}
	return power
}

func insert(boxes [][]lens, label string, focallength int) {
	boxNum := hash(label)
	for i, lens := range boxes[boxNum] {
		if lens.label == label {
			boxes[boxNum][i].focallength = focallength
			return
		}

	}
	boxes[boxNum] = append(boxes[boxNum], lens{label, focallength})
}

func remove(boxes [][]lens, label string) {
	boxNum := hash(label)
	for i, lens := range boxes[boxNum] {
		if lens.label == label {
			boxes[boxNum] = append(boxes[boxNum][:i], boxes[boxNum][i+1:]...)
			return
		}
	}
}

func calculateSteps(steps []string) {
	boxes := make([][]lens, 256)

	for _, step := range steps {
		if index := strings.Index(step, "="); index != -1 {
			label := step[:index]
			amount := int(step[index+1] - '0')

			insert(boxes, label, amount)
		}
		if index := strings.Index(step, "-"); index != -1 {
			label := step[:index]
			remove(boxes, label)
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
	for scanner.Scan() {
		line := scanner.Text()
		steps = strings.Split(line, ",")
	}
	calculateSteps(steps)

}
