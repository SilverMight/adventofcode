package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	fiveOfAKind  = 6
	fourOfAKind  = 5
	fullHouse    = 4
	threeOfaKind = 3
	twoPair      = 2
	onePair      = 1
	highCard     = 0
)

func compareCards(card1, card2 string) bool {
	ranks := "AKQT98765432J"

	for i := 0; i < 5; i++ {
		if card1[i] != card2[i] {
			index1 := strings.Index(ranks, string(card1[i]))
			index2 := strings.Index(ranks, string(card2[i]))

			return index1 > index2
		}
	}

	return false
}
func getRank(hand string) int {
	cardMap := make(map[rune]int)

	for _, card := range hand {
		cardMap[card]++
	}
	fmt.Println(hand, cardMap)

	numPairs := 0
	numTriplets := 0
	numJokers, hasJoker := cardMap['J']

	for card, amount := range cardMap {
		if amount == 5 {
			return fiveOfAKind
		}
		if amount == 4 && card != 'J' {
			if hasJoker && numJokers == 1 {
				return fiveOfAKind
			}
			return fourOfAKind
		}
		if amount == 2 {
			numPairs++
		}
		if amount == 3 {
			numTriplets++
		}

	}
	if hasJoker && numJokers == 4 {
		return fiveOfAKind
	}

	if numPairs == 1 && hasJoker && numJokers == 3 {
		return fiveOfAKind
	}
	if numTriplets == 1 && cardMap['J'] != 3 {
		if hasJoker && numJokers == 1 {
			return fourOfAKind
		}
		if hasJoker && numJokers == 2 {
			return fiveOfAKind
		}
		if numPairs == 1 {
			return fullHouse
		}
		return threeOfaKind
	}
	if numPairs == 2 {
		if hasJoker && numJokers == 2 {
			return fourOfAKind
		}
		if hasJoker && numJokers == 1 {
			return fullHouse
		}
		return twoPair
	}
	if numPairs == 1 && cardMap['J'] != 2 {
		if hasJoker && numJokers == 1 {
			return threeOfaKind
		}
		if hasJoker && numJokers == 2 {
			return fourOfAKind
		}
		if hasJoker && numJokers == 3 {
			return fullHouse
		}
		return onePair
	}

	// lone jokers case
	if hasJoker && numJokers == 1 {
		return onePair
	}
	if hasJoker && numJokers == 2 {
		return threeOfaKind
	}
	if hasJoker && numJokers == 3 {
		return fourOfAKind
	}

	return highCard
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	cardsToBids := make(map[string]int)

	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Fields(line)

		cardsToBids[fields[0]], _ = strconv.Atoi(fields[1])
	}

	// seven ranks
	rankMap := make([][]string, 7)
	for card := range cardsToBids {
		rank := getRank(card)
		rankMap[rank] = append(rankMap[rank], card)
	}

	// Tiebreaking
	for rank := range rankMap {
		sort.SliceStable(rankMap[rank], func(i, j int) bool {
			return compareCards(rankMap[rank][i], rankMap[rank][j])
		})
	}

	winnings := 0
	rankNum := 1
	for _, cards := range rankMap {
		for _, card := range cards {
			winnings += cardsToBids[card] * rankNum
			rankNum++
		}
	}

	fmt.Println(winnings)

}
