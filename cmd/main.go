package main

import (
	"advent-of-code-2023/pkg/solution"
	"advent-of-code-2023/pkg/util"
	"fmt"
	"log"
)

const ioFilepathPrefix = "./assets/day_"

func main() {
	var dayNumber, taskNumber int

	fmt.Println("Enter the day number (from 1 to 9)")
	_, err := fmt.Scanf("%d", &dayNumber)
	if err != nil || !(dayNumber >= 1 && dayNumber <= 25) {
		log.Fatal("Incorrect day number")
	}

	fmt.Println("Enter task number (1 or 2)")
	_, err = fmt.Scanf("%d", &taskNumber)
	if err != nil || !(taskNumber == 1 || taskNumber == 2) {
		log.Fatal("Incorrect task number")
	}

	inputFilepath := fmt.Sprintf("%s%d/input.txt", ioFilepathPrefix, dayNumber)
	lines := util.ReadLines(inputFilepath)

	var answer string
	switch dayNumber {
	case 1:
		if taskNumber == 1 {
			answer = solution.Solve_1_1(lines)
		} else {
			answer = solution.Solve_1_2(lines)
		}

	case 2:
		if taskNumber == 1 {
			answer = solution.Solve_2_1(lines)
		} else {
			answer = solution.Solve_2_2(lines)
		}

	case 3:
		if taskNumber == 1 {
			answer = solution.Solve_3_1(lines)
		} else {
			answer = solution.Solve_3_2(lines)
		}

	case 4:
		if taskNumber == 1 {
			answer = solution.Solve_4_1(lines)
		} else {
			answer = solution.Solve_4_2(lines)
		}
	case 5:
		if taskNumber == 1 {
			answer = solution.Solve_5_1(lines)
		} else {
			answer = solution.Solve_5_2(lines)
		}
	case 6:
		if taskNumber == 1 {
			answer = solution.Solve_6_1(lines)
		} else {
			answer = solution.Solve_6_2(lines)
		}
	case 7:
		if taskNumber == 1 {
			answer = solution.Solve_7_1(lines)
		} else {
			answer = solution.Solve_7_2(lines)
		}
	case 8:
		if taskNumber == 1 {
			answer = solution.Solve_8_1(lines)
		} else {
			answer = solution.Solve_8_2(lines)
		}
	case 9:
		if taskNumber == 1 {
			answer = solution.Solve_9_1(lines)
		} else {
			answer = solution.Solve_9_2(lines)
		}
	default:
		log.Fatalf("Tasks from day %d are not implemented yet", dayNumber)
	}

	outputFilepath := fmt.Sprintf("%s%d/output_%d.txt", ioFilepathPrefix, dayNumber, taskNumber)
	util.WriteString(outputFilepath, answer)
}
