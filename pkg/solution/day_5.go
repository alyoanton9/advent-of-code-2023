package solution

import (
	"advent-of-code-2023/pkg/util"
	"fmt"
	"math"
	"strings"
)

func Solve_5_1(lines []string) string {
	seedToLocation := getMap(lines)
	answer := calcMinLocation(seedToLocation)

	return fmt.Sprint(answer)
}

func Solve_5_2(lines []string) string {
	seedToLocation := getMap(lines)
	answer := calcMinLocationFromSeedRange(seedToLocation)

	return fmt.Sprint(answer)
}

func calcMinLocation(seedToLocation SeedToLocationMap) int {
	var minLocation = math.MaxInt

	for _, seed := range seedToLocation.seeds {
		value := seed
		for _, mapping := range seedToLocation.mappings {
			for _, mapRow := range mapping {
				if value >= mapRow.source && value < mapRow.source+mapRow.delta {
					value += mapRow.destination - mapRow.source
					break
				}
			}
		}
		location := value
		minLocation = min(minLocation, location)
	}

	return minLocation
}

func calcMinLocationFromSeedRange(seedToLocation SeedToLocationMap) int {
	var minLocation = math.MaxInt

	for ind := 0; ind < len(seedToLocation.seeds); ind += 2 {
		startSeed := seedToLocation.seeds[ind]
		endSeed := seedToLocation.seeds[ind] + seedToLocation.seeds[ind+1]
		for seed := startSeed; seed < endSeed; seed++ {
			// copy-paste from original function
			value := seed
			for _, mapping := range seedToLocation.mappings {
				for _, mapRow := range mapping {
					if value >= mapRow.source && value < mapRow.source+mapRow.delta {
						value += mapRow.destination - mapRow.source
						break
					}
				}
			}
			location := value
			minLocation = min(minLocation, location)
		}
	}

	return minLocation
}

type SeedToLocationMap struct {
	seeds    []int
	mappings [][]MapRow
}

type MapRow struct {
	source      int
	destination int
	delta       int
}

func getMap(lines []string) SeedToLocationMap {
	seedToLocation := SeedToLocationMap{}

	seeds := util.GetAllIntsFromString(lines[0])
	seedToLocation.seeds = seeds

	ind := 2
	for ind < len(lines) {
		if strings.HasSuffix(lines[ind], "map:") {
			ind++

			mapRows := make([]MapRow, 0)
			for ind < len(lines) && lines[ind] != "" {
				numbers := util.GetAllIntsFromString(lines[ind])
				mapRows = append(mapRows, MapRow{source: numbers[1], destination: numbers[0], delta: numbers[2]})
				ind++
			}
			seedToLocation.mappings = append(seedToLocation.mappings, mapRows)
		}
		ind++
	}

	return seedToLocation
}
