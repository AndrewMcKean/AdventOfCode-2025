package main

import (
	"bufio"
	"os"
	"sort"
	"strings"
	"strconv"
)

func main() {
	fi, _ := os.Open("inputTboxData.txt")
	scanner := bufio.NewScanner(fi)

	var ranges [][2]int

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, "-")
		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		ranges = append(ranges, [2]int{start, end})
	}

	// Sort by start for easier overlap checking
	sort.Slice(ranges, func(i int, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	assessedRanges := make([][2]int, 0)

	currentRange := ranges[0]

	for i := 1; i < len(ranges); i++ {
		rangeToAssessStart := ranges[i][0]
		currentRangeEnd := currentRange[1] + 1

		if rangeToAssessStart <= currentRangeEnd {
			rangeToAssessEnd := ranges[i][1]

			if rangeToAssessEnd >= currentRangeEnd {
				currentRange[1] = ranges[i][1]
			} 
		} else {
			// No overlap, push current into 
			assessedRanges = append(assessedRanges, currentRange)
			currentRange = ranges[i]
		}

	}

	assessedRanges = append(assessedRanges, currentRange)

	// Count
	numValidIDs := 0
	for _,  assessedRange := range assessedRanges {
		// Add 1 as it's inclusive
		numValidIDs += (assessedRange[1] + 1) - assessedRange[0]
	}

	println("Total numValidIDs: ", numValidIDs)

}
