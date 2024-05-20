package solution

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Solve_2_1(lines []string) string {
	games := getGameInfo(lines)

	maxRed, maxGreen, maxBlue := 12, 13, 14
	answer := calcIdsSum(games, maxRed, maxGreen, maxBlue)

	return fmt.Sprint(answer)
}

func Solve_2_2(lines []string) string {
	games := getGameInfo(lines)

	answer := calcPowerSum(games)

	return fmt.Sprint(answer)
}

const (
	red   = "red"
	green = "green"
	blue  = "blue"
)

type GameInfo struct {
	id           int
	colorNumbers []ColorNumber
}

type ColorNumber struct {
	number int
	color  string
}

func getGameInfo(lines []string) []GameInfo {
	games := make([]GameInfo, 0)

	for _, line := range lines {
		gameAndSubsets := strings.Split(line, ": ")
		gameIdStr := strings.Split(gameAndSubsets[0], " ")[1]
		subsets := strings.Split(gameAndSubsets[1], " ")

		colorNumbers := make([]ColorNumber, 0)
		for ind := 0; ind < len(subsets); ind += 2 {
			num, err := strconv.Atoi(subsets[ind])
			if err != nil {
				log.Fatalf("error parsing cube number: %s", err.Error())
			}

			color := strings.TrimSuffix(subsets[ind+1], ",")
			color = strings.TrimSuffix(color, ";")

			colorNumbers = append(colorNumbers, ColorNumber{number: num, color: color})
		}

		gameId, err := strconv.Atoi(gameIdStr)
		if err != nil {
			log.Fatalf("error parsing gameId: %s", err.Error())
		}

		games = append(games, GameInfo{id: gameId, colorNumbers: colorNumbers})
	}

	return games
}

func calcPowerSum(games []GameInfo) int {
	sum := 0

	for _, game := range games {

		maxRed, maxGreen, maxBlue := 0, 0, 0
		for _, colorNumber := range game.colorNumbers {
			switch colorNumber.color {
			case red:
				maxRed = max(colorNumber.number, maxRed)
			case green:
				maxGreen = max(colorNumber.number, maxGreen)
			case blue:
				maxBlue = max(colorNumber.number, maxBlue)
			}
		}

		power := maxRed * maxGreen * maxBlue
		sum += power
	}

	return sum
}

func calcIdsSum(games []GameInfo, maxRed int, maxGreen int, maxBlue int) int {
	sum := 0

	for _, game := range games {

		possible := true
		for _, colorNumber := range game.colorNumbers {
			if colorNumber.number > maxRed && colorNumber.color == red ||
				colorNumber.number > maxGreen && colorNumber.color == green ||
				colorNumber.number > maxBlue && colorNumber.color == blue {
				possible = false
				break
			}
		}

		if possible {
			sum += game.id
		}
	}

	return sum
}
