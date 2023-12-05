package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getNumMatches(winningNums, guessedNums []string) int {
    count := 0
    for _, guessString := range guessedNums {
        guess, _ := strconv.Atoi(guessString)
        for _, winString := range winningNums {
            win, _ := strconv.Atoi(winString)

            if win == guess {
                count++
            }
        }

    }

    return count
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    sum := 0
    for scanner.Scan() {
        line := scanner.Text()

        // Not interested in the game number
        line = line[strings.Index(line, ":") + 1:]


        parts := strings.Split(line, "|")
        winningNums := strings.Fields(parts[0])
        guessedNums := strings.Fields(parts[1])
        

        numCorrectGuesses := getNumMatches(winningNums, guessedNums)

        if numCorrectGuesses > 0 {
            // 2 ^ our guesses
            sum += (1 << (numCorrectGuesses - 1))
        }
    }

    fmt.Println(sum) 
}
