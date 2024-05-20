package solution

import (
	"advent-of-code-2023/pkg/util"
	"fmt"
	"github.com/samber/lo"
	"regexp"
	"strings"
	"sync"
)

func Solve_8_1(lines []string) string {
	instructions := parseInstructions(lines)
	answer := calcSteps(instructions)

	return fmt.Sprint(answer)
}

func Solve_8_2(lines []string) string {
	instructions := parseInstructions(lines)
	answer := calcStepsParallel(instructions)

	return fmt.Sprint(answer)
}

type Step struct {
	left  string
	right string
}

type Instructions struct {
	sequence string
	steps    map[string]Step
}

func calcSteps(instructions Instructions) int {
	node := "AAA"

	directions := []rune(instructions.sequence)
	seqLen := len(directions)

	step := 0
	for node != "ZZZ" {
		direction := directions[step%seqLen]

		if direction == 'L' {
			node = instructions.steps[node].left
		} else {
			node = instructions.steps[node].right
		}

		step++
	}

	return step
}

func calcStepsParallel(instructions Instructions) int64 {
	nodes := util.Keys(instructions.steps)
	nodes = lo.Filter(nodes, func(node string, _ int) bool {
		return strings.HasSuffix(node, "A")
	})

	directions := []rune(instructions.sequence)
	seqLen := int64(len(directions))

	steps := make([]int64, 0)

	// parallel calculations just for fun
	wg := sync.WaitGroup{}
	for _, node := range nodes {
		wg.Add(1)

		go func(node string) {
			defer wg.Done()

			var step int64 = 0
			for !strings.HasSuffix(node, "Z") {
				ind := step % seqLen
				direction := directions[ind]

				if direction == 'L' {
					node = instructions.steps[node].left
				} else {
					node = instructions.steps[node].right
				}

				step++
			}

			steps = append(steps, step)
		}(node)
	}
	wg.Wait()

	stepsLCM := util.LCM(
		steps[0], steps[1], steps[2], steps[3], steps[4], steps[5])

	return stepsLCM
}

func parseInstructions(lines []string) Instructions {
	sequence := lines[0]
	steps := make(map[string]Step)

	for _, line := range lines[2:] {
		reMatchSteps := regexp.MustCompile(`([A-Z]+)`)
		nodes := reMatchSteps.FindAllString(line, -1)

		steps[nodes[0]] = Step{left: nodes[1], right: nodes[2]}
	}

	return Instructions{sequence: sequence, steps: steps}
}
