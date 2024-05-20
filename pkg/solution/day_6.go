package solution

import (
	"advent-of-code-2023/pkg/util"
	"fmt"
	"math"
)

func Solve_6_1(lines []string) string {
	timeDistances := getTimeDistances(lines)
	answer := calcNumberOfWays(timeDistances)

	return fmt.Sprint(answer)
}

func Solve_6_2(lines []string) string {
	timeDistance := getGluedTimeDistance(lines)
	answer := calcNumberOfWays(timeDistance)

	return fmt.Sprint(answer)
}

type TimeDistance struct {
	Time     int
	Distance int
}

func getGluedTimeDistance(lines []string) []TimeDistance {
	time := util.GlueAllIntsInString(lines[0])
	dist := util.GlueAllIntsInString(lines[1])

	return []TimeDistance{{Time: time, Distance: dist}}
}

func getTimeDistances(lines []string) []TimeDistance {
	times := util.GetAllIntsFromString(lines[0])
	dists := util.GetAllIntsFromString(lines[1])

	timeDistances := make([]TimeDistance, 0)
	for i := 0; i < len(times); i++ {
		timeDistances = append(timeDistances, TimeDistance{Time: times[i], Distance: dists[i]})
	}

	return timeDistances
}

func calcNumberOfWays(timeDistances []TimeDistance) int {
	/*
		speed * move_time > dist
		speed = x
		move_time = time - x
		x * (time - x) > dist
		x^2 - time*x + dist < 0

		x1 = (time + sqrt(time^2 - 4*dist))/2
		x2 = (time - sqrt(time^2 - 4*dist))/2

		x^2 - time*x + dist < 0 for x=x1..x2
	*/

	waysNumber := 1

	for _, timeDistance := range timeDistances {
		time := timeDistance.Time
		dist := timeDistance.Distance
		solution1, solution2 := solveQuadraticEquation(1, -float64(time), float64(dist))

		minSpeed := int(math.Ceil(solution1))
		maxSpeed := int(math.Floor(solution2))

		waysNumber *= maxSpeed - minSpeed + 1
	}

	return waysNumber
}

func solveQuadraticEquation(a, b, c float64) (float64, float64) {
	discriminant := b*b - 4*a*c
	discriminantSqrt := math.Sqrt(discriminant)

	solution1 := (-b - discriminantSqrt) / (2 * a)
	solution2 := (-b + discriminantSqrt) / (2 * a)

	return solution1, solution2
}
