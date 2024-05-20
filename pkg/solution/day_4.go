package solution

import (
	"advent-of-code-2023/pkg/util"
	"fmt"
	"strings"
)

func Solve_4_1(lines []string) string {
	cards := getCards(lines)
	answer := calcCardPoints(cards)

	return fmt.Sprint(answer)
}

func Solve_4_2(lines []string) string {
	cards := getCards(lines)
	answer := calcCardNumber(cards)

	return fmt.Sprint(answer)
}

type Card struct {
	index          int
	winningNumbers map[int]struct{}
	cardNumbers    []int
}

func getCards(lines []string) []Card {
	cards := make([]Card, 0)

	for ind, line := range lines {
		numbersSplit := strings.Split(line, "|")

		winningNumbers := util.GetAllIntsFromString(numbersSplit[0])[1:]
		winningNumbersSet := util.SliceToSet(winningNumbers)

		cardNumbers := util.GetAllIntsFromString(numbersSplit[1])

		cards = append(cards, Card{
			index:          ind + 1,
			winningNumbers: winningNumbersSet,
			cardNumbers:    cardNumbers,
		})
	}

	return cards
}

func calcCardPoints(cards []Card) int {
	sumPoints := 0

	for _, card := range cards {
		cardPoint := 1
		for _, cardNumber := range card.cardNumbers {
			if _, ok := card.winningNumbers[cardNumber]; ok {
				cardPoint *= 2
			}
		}

		sumPoints += cardPoint / 2
	}

	return sumPoints
}

func calcCardNumber(cards []Card) int {
	copies := make(map[int]int)
	for _, card := range cards {
		copies[card.index] = 1
	}

	sumCards := 0

	for _, card := range cards {
		matchingNum := 0
		for _, cardNumber := range card.cardNumbers {
			if _, ok := card.winningNumbers[cardNumber]; ok {
				matchingNum += 1
			}
		}

		sumCards += copies[card.index]

		for i := 1; card.index+i <= len(cards) && i <= matchingNum; i++ {
			copies[card.index+i] += copies[card.index]
		}
	}

	return sumCards
}
