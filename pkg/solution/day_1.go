package solution

import (
	"fmt"
	"regexp"
	"strconv"
)

func Solve_1_1(lines []string) string {
	answer := calcCalibrationValue(lines, false)
	return fmt.Sprint(answer)
}

func Solve_1_2(lines []string) string {
	answer := calcCalibrationValue(lines, true)
	return fmt.Sprint(answer)
}

func calcCalibrationValue(lines []string, spelling bool) int {
	digitRegexp := `\d`
	if spelling {
		digitRegexp += `|one|two|three|four|five|six|seven|eight|nine`
	}

	reFirstDigit := regexp.MustCompile(fmt.Sprintf(`[a-z]*?(%s)[a-z0-9]*`, digitRegexp))
	reLastDigit := regexp.MustCompile(fmt.Sprintf(`[a-z0-9]*(%s)[a-z]*`, digitRegexp))

	sum := 0

	for _, line := range lines {
		firstDigitStr := reFirstDigit.FindStringSubmatch(line)[1]
		lastDigitStr := reLastDigit.FindStringSubmatch(line)[1]
		number := stringToNumber(firstDigitStr)*10 + stringToNumber(lastDigitStr)

		sum += number
	}

	return sum
}

func stringToNumber(s string) int {
	var number int

	switch s {
	case "one":
		number = 1
	case "two":
		number = 2
	case "three":
		number = 3
	case "four":
		number = 4
	case "five":
		number = 5
	case "six":
		number = 6
	case "seven":
		number = 7
	case "eight":
		number = 8
	case "nine":
		number = 9
	default:
		number, _ = strconv.Atoi(s)
	}

	return number
}
