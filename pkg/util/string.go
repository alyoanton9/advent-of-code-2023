package util

import (
	"github.com/samber/lo"
	"regexp"
	"strconv"
)

var reMatchNumbers = regexp.MustCompile(`-?\d+`)

func GetAllIntsFromString(str string) []int {
	numbersStr := reMatchNumbers.FindAllString(str, -1)
	numbers := lo.Map(numbersStr, func(numberStr string, _ int) int {
		num, _ := strconv.Atoi(numberStr)
		return num
	})

	return numbers
}

func GlueAllIntsInString(str string) int {
	numbersStr := reMatchNumbers.FindAllString(str, -1)
	numberStr := lo.Reduce(numbersStr, func(agg string, item string, _ int) string {
		return agg + item
	}, "")

	number, _ := strconv.Atoi(numberStr)

	return number
}
