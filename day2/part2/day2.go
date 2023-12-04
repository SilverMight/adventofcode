package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)



func main() {
    scanner := bufio.NewScanner(os.Stdin)

    sumOfPowers := 0

    for scanner.Scan() {
        line := scanner.Text()

        // Don't care about semicolons
        line = strings.ReplaceAll(line, ";", ",")

        // I don't really want the game number
        games := strings.Split(line[strings.Index(line, ":") + 1:], ",")

        // Hold the largest number for each color we find
        minNumberOfCubesPerColor := make(map[string]int)

        for _, game := range games {
            amountAndColor := strings.Split(strings.TrimSpace(game), " ")

            amount, _:= strconv.Atoi(amountAndColor[0])

            // If amount is larger than what we've seen for this color before, new largest number 
            if(amount > minNumberOfCubesPerColor[amountAndColor[1]]) {
                minNumberOfCubesPerColor[amountAndColor[1]] = amount
            }
        }

        // calculate power of a set of cubes
        power := 1
        for _, amount := range minNumberOfCubesPerColor {
            power *= amount
        }

        sumOfPowers += power
    }

    fmt.Println(sumOfPowers)
}
