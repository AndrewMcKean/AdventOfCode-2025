package main

import (
	"bufio"
	"os"
	"strings"
	"strconv"
)


func main() {
	fi, _ := os.Open("inputTbox.txt")
	scanner := bufio.NewScanner(fi)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fiData, _ := os.Open("inputTboxData.txt")
	dataLines := []string{}
	dataScanner := bufio.NewScanner(fiData)
	for dataScanner.Scan() {
		dataLines = append(dataLines, dataScanner.Text())
	}

	println("Getting validNums")
	// Create data reference
	freshIDs := make(map[int]int)
	for _, dataLine := range dataLines {
		parts := strings.Split(dataLine, "-")

		start, _ := strconv.Atoi(parts[0])
		end, _ := strconv.Atoi(parts[1])

		for _, ID := range lines {
			numID, _ := strconv.Atoi(ID)
			numExists := numID >= start && numID <= end

			if numExists {
				freshIDs[numID] = 1	
			}
		
		}
	}

	println("Num Fresh IDs: ", len(freshIDs))
}

