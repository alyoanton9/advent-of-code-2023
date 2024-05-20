package util

import (
	"bufio"
	"github.com/samber/lo"
	"log"
	"os"
)

func LinesToRunes(lines []string) [][]rune {
	runes := lo.Map(lines, func(line string, _ int) []rune {
		return []rune(line)
	})

	return runes
}

func ReadLines(filepath string) []string {
	inputFile, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func WriteString(filepath string, content string) {
	outputFile, err := os.Create(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)

	_, err = writer.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}

	writer.Flush()
}
