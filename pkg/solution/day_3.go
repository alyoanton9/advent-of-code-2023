package solution

import (
	"advent-of-code-2023/pkg/util"
	"fmt"
	"strconv"
	"unicode"
)

func Solve_3_1(lines []string) string {
	runes := util.LinesToRunes(lines)
	answer, _ := calcPartNumbersSum(runes, isSymbol)

	return fmt.Sprint(answer)
}

func Solve_3_2(lines []string) string {
	runes := util.LinesToRunes(lines)
	_, answer := calcPartNumbersSum(runes, isGear)

	return fmt.Sprint(answer)
}

type predicate func(rune) bool

func calcPartNumbersSum(schematic [][]rune, symbolFunc predicate) (int, int) {
	maxRow := len(schematic)
	maxCol := len(schematic[0])

	row, col := 0, 0
	sumPartNumbers, sumGearRatios := 0, 0

	gears := make(map[string][]int)

	for row < maxRow {

		col = 0
		for col < maxCol {
			isPartNumber := false
			number := 0

			gearsCoord := make([][]int, 0)

			if unicode.IsDigit(schematic[row][col]) {
				digitRunes := make([]rune, 0)

				// left adjacent symbols
				for r := row - 1; r <= row+1; r++ {
					if r >= 0 && r < maxRow && col > 0 {
						if symbolFunc(schematic[r][col-1]) {
							isPartNumber = true
							gearsCoord = append(gearsCoord, []int{r, col - 1})
						}
					}
				}

				// upper and lower adjacent symbols
				for col < maxCol && unicode.IsDigit(schematic[row][col]) {
					if row > 0 && symbolFunc(schematic[row-1][col]) {
						isPartNumber = true
						gearsCoord = append(gearsCoord, []int{row - 1, col})
					}

					if row < maxRow-1 && symbolFunc(schematic[row+1][col]) {
						isPartNumber = true
						gearsCoord = append(gearsCoord, []int{row + 1, col})
					}
					digitRunes = append(digitRunes, schematic[row][col])
					col++
				}

				// right adjacent symbols
				for r := row - 1; r <= row+1; r++ {
					if r >= 0 && r < maxRow && col < maxCol {
						if symbolFunc(schematic[r][col]) {
							isPartNumber = true
							gearsCoord = append(gearsCoord, []int{r, col})
						}
					}
				}

				if isPartNumber {
					number, _ = strconv.Atoi(string(digitRunes))
					sumPartNumbers += number

					for _, gearCoord := range gearsCoord {
						key := fmt.Sprintf("%d:%d", gearCoord[0], gearCoord[1])
						gears[key] = append(gears[key], number)
					}
				}
			}
			col++
		}
		row++
	}

	for _, partNumbers := range gears {
		if len(partNumbers) == 2 {
			sumGearRatios += partNumbers[0] * partNumbers[1]
		}
	}

	return sumPartNumbers, sumGearRatios
}

func isSymbol(r rune) bool {
	return !unicode.IsDigit(r) && r != '.'
}

func isGear(r rune) bool {
	return r == '*'
}
