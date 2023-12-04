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

    maxCubesPerColor := map[string]int {
        "red" : 12,
        "green": 13,
        "blue": 14,
    }

    gameId := 1
    sum := 0

    for scanner.Scan() {
        line := scanner.Text()

        // Don't care about semicolons
        line = strings.ReplaceAll(line, ";", ",")



        // I don't really want the game number
        games := strings.Split(line[strings.Index(line, ":") + 1:], ",")

        possibleRound := true

        for _, game := range games {
            amountAndColor := strings.Split(strings.TrimSpace(game), " ")

            amount, _:= strconv.Atoi(amountAndColor[0])

            if(amount > maxCubesPerColor[amountAndColor[1]]) {
                possibleRound = false
            }
        }

        if(possibleRound) {
            sum += gameId
        }

        gameId++

    }

    fmt.Println(sum)
}
