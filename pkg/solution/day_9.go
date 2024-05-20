package solution

import (
	"advent-of-code-2023/pkg/util"
	"fmt"
	"github.com/samber/lo"
	"sync"
)

func Solve_9_1(lines []string) string {
	answer := calcExtrapolations(lines, false)

	return fmt.Sprint(answer)
}

func Solve_9_2(lines []string) string {
	answer := calcExtrapolations(lines, true)

	return fmt.Sprint(answer)
}

func calcExtrapolations(lines []string, backwards bool) int {
	sum := 0

	wg := sync.WaitGroup{}
	for _, line := range lines {
		wg.Add(1)

		go func(line string) {
			defer wg.Done()

			values := util.GetAllIntsFromString(line)

			deltas := make([]int, 0)

			delta := values[len(values)-1]
			if backwards {
				delta = values[0]
			}
			deltas = append(deltas, delta)

			sign := -1
			for lo.Count(values, 0) != len(values) {
				newValues := make([]int, 0)
				for i := 0; i < len(values)-1; i++ {
					newValues = append(newValues, values[i+1]-values[i])
				}

				values = newValues

				var nextDelta int
				if backwards {
					nextDelta = sign * values[0]
				} else {
					nextDelta = values[len(values)-1]
				}

				deltas = append(deltas, nextDelta)
				sign *= -1
			}

			sum += lo.Sum(deltas)
		}(line)
	}
	wg.Wait()

	return sum
}
