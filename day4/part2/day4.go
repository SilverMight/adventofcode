package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getNumOfMatches(winningNums, guessedNums []string) int {
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

    numMatchesPerGame := []int{}
    sum := 0
    for scanner.Scan() {
        line := scanner.Text()

        // Not interested in the game number
        line = line[strings.Index(line, ":") + 1:]


        parts := strings.Split(line, "|")
        winningNums := strings.Fields(parts[0])
        guessedNums := strings.Fields(parts[1])
        
        numMatchesPerGame = append(numMatchesPerGame, getNumOfMatches(winningNums, guessedNums))
    }

    
    amountOfCards := make([]int, len(numMatchesPerGame))

    for game, numMatches := range numMatchesPerGame {
        // We have implicitly have one card 
        amountOfCards[game]++
        
        // For the amount of matches we made, add the amount of current cards we have
        for index := game + 1; index <= game + numMatches && index < len(numMatchesPerGame); index++ {
            amountOfCards[index] += amountOfCards[game] 
        }

    }

    // Accumulate
    for _, amount := range amountOfCards {
        sum += amount
    }

    fmt.Println(sum) 
}
