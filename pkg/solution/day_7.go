package solution

import (
	"advent-of-code-2023/pkg/util"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func Solve_7_1(lines []string) string {
	handCards := parseHandCards(lines)
	answer := calcTotalWinnings(handCards, false)

	return fmt.Sprint(answer)
}

func Solve_7_2(lines []string) string {
	handCards := parseHandCards(lines)
	answer := calcTotalWinnings(handCards, true)

	return fmt.Sprint(answer)
}

type HandCard struct {
	hand string
	bid  int
}

func calcTotalWinnings(handCards []HandCard, withJoker bool) int {
	sort.Slice(handCards, func(i, j int) bool {
		iRank, jRank := getKindRank(handCards[i].hand, withJoker), getKindRank(handCards[j].hand, withJoker)
		if iRank == jRank {
			return less([]rune(handCards[i].hand), []rune(handCards[j].hand), withJoker)
		} else {
			return iRank < jRank
		}
	})

	totalWinnings := 0
	for i, handCard := range handCards {
		totalWinnings += (i + 1) * handCard.bid
	}

	return totalWinnings
}

func getKindRank(hand string, withJoker bool) string {
	counter := util.Count([]rune(hand))

	freqs := make([]int, 0)
	if withJoker {
		freqs = calcFreqsWithJokers(counter)
	} else {
		for _, freq := range counter {
			freqs = append(freqs, freq)
		}
		sort.Ints(freqs)
	}

	rank := ""
	for i := len(freqs) - 1; i >= 0; i-- {
		rank += fmt.Sprint(freqs[i])
	}

	return rank
}

func less(x []rune, y []rune, withJoker bool) bool {
	order := []rune{'A', 'K', 'Q'}
	if withJoker {
		order = append(order, 'T')
	} else {
		order = append(order, 'J', 'T')
	}

	for i := '9'; i > '1'; i-- {
		order = append(order, i)
	}
	if withJoker {
		order = append(order, 'J')
	}

	for i := 0; i < len(x); i++ {
		indexX := slices.Index(order, x[i])
		indexY := slices.Index(order, y[i])

		if indexX != indexY {
			return indexX > indexY
		}
	}

	return false
}

func calcFreqsWithJokers(counter map[rune]int) []int {
	five := []int{5}

	jokersNumber := counter['J']

	freqs := make([]int, 0)
	for r, freq := range counter {
		if r != 'J' {
			freqs = append(freqs, freq)
		}
	}
	sort.Ints(freqs)

	size := len(freqs)
	if size == 0 {
		return five
	}

	freqs[size-1] += jokersNumber

	return freqs
}

func parseHandCards(lines []string) []HandCard {
	handCards := make([]HandCard, 0)

	for _, line := range lines {
		handSplit := strings.Split(line, " ")
		bid, _ := strconv.Atoi(handSplit[1])
		handCards = append(handCards, HandCard{hand: handSplit[0], bid: bid})
	}

	return handCards
}
