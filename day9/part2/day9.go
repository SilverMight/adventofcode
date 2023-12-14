package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculateHistory(array []int) [][]int {
	differences := [][]int{}

	differences = append(differences, array)

	allZeroes := false
	row := 0
	for !allZeroes {
		differences = append(differences, []int{})
		for i := 0; i < len(differences[row])-1; i++ {
			difference := differences[row][i+1] - differences[row][i]

			differences[row+1] = append(differences[row+1], difference)
		}

		allZeroes = true
		for _, value := range differences[row+1] {
			if value != 0 {
				allZeroes = false
				break
			}
		}

		row++
	}
	return differences

}

func extrapolate(differences [][]int) int {
	length := len(differences)

	sum := 0
	for i := length - 1; i >= 1; i-- {
		sum += differences[i-1][len(differences[i-1])-1]
	}

	return sum
}

func stringsToIntegers(array []string) []int {
	ret := make([]int, len(array))
	for index, number := range array {
		ret[index], _ = strconv.Atoi(number)
	}

	return ret
}

func reverseArray(array []int) {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	sum := 0
	for scanner.Scan() {
		history := scanner.Text()
		historyFields := strings.Fields(history)
		historyNumbers := stringsToIntegers(historyFields)
		reverseArray(historyNumbers)

		fmt.Println(calculateHistory(historyNumbers))
		sum += extrapolate(calculateHistory(historyNumbers))
	}
	fmt.Println(sum)
}
