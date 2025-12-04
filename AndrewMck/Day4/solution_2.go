package main

import (
	"bufio"
	"os"
)

const SEARCH_CHAR = "@"

func getIndicesToCheck(currLineIndex int, currentCharIndex int, lengthOfLine int, totalLines int) []([]int) {
	returnSlice := []([]int){}

	// Prev indices
	prevIndices := []int{}
	
	if currLineIndex > 0 {
		if currentCharIndex > 0 {
			prevIndices = append(prevIndices, currentCharIndex - 1)
		} else {
			prevIndices = append(prevIndices, -1)
		}

		prevIndices = append(prevIndices, currentCharIndex)

		if currentCharIndex < lengthOfLine - 1 {
			prevIndices = append(prevIndices, currentCharIndex + 1)
		} else {
			prevIndices = append(prevIndices, -1)
		}
	}

	returnSlice = append(returnSlice, prevIndices)

	// Curr indices
	currIndices := []int{}
	if currentCharIndex > 0 {
		currIndices = append(currIndices, currentCharIndex - 1) 
	} else {
		currIndices = append(currIndices, -1)
	}

	if currentCharIndex < lengthOfLine - 1 {
		currIndices = append(currIndices, currentCharIndex + 1)
	} else {
		currIndices = append(currIndices, -1)
	}

	returnSlice = append(returnSlice, currIndices)


	// Next line indices
	nextIndices := []int{}
	if currLineIndex < totalLines - 1 { 
		if currentCharIndex > 0 {
			nextIndices = append(nextIndices, currentCharIndex - 1)
		} else {
			nextIndices = append(nextIndices, -1)
		}

		nextIndices = append(nextIndices, currentCharIndex)

		if currentCharIndex < lengthOfLine - 1 {
			nextIndices = append(nextIndices, currentCharIndex + 1)
		} else {
			nextIndices = append(nextIndices, -1)
		}
	}

	returnSlice = append(returnSlice, nextIndices)

	return returnSlice
}

func main() {
	fi, _ := os.Open("inputTbox.txt")
	scanner := bufio.NewScanner(fi)

	lines := []string{}
 
	totalRolls := 0

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	goAgain := true

	for goAgain {
  	numLinesChanged := 0

	for i := 0; i < len(lines); i++ {
		prevLine := "" 
		if i > 0 {
			prevLine = lines[i-1]
		}
		
		currLine := lines[i]

		nextLine := ""
		if i < len(lines) - 1 {
			nextLine = lines[i+1]
		}

		println("Assessing row: ", i)

		updatedLine := ""

		for j := 0; j < len(currLine); j++ {
			if string(currLine[j]) != SEARCH_CHAR {
				updatedLine += string(currLine[j])
				continue
			}

			indicesToCheck := getIndicesToCheck(i, j, len(currLine), len(lines))
			numNearbyRolls := 0
			
			// Check prevLine
			for _, x := range indicesToCheck[0] {
				if x >= 0 {
					if string(prevLine[x]) == SEARCH_CHAR {
						numNearbyRolls++
					}
				}
			}

			// Check currLine
			for _, x := range indicesToCheck[1] {
				if x >= 0 {
					if string(currLine[x]) == SEARCH_CHAR {
						numNearbyRolls++
					}
				}
			}

			// Check nextLine
			for _, x := range indicesToCheck[2] {
				if x >= 0 {
					if string(nextLine[x]) == SEARCH_CHAR {
						numNearbyRolls++
					}
				}
			}

			if numNearbyRolls < 4 {
				totalRolls++
				updatedLine += string("x")
			} else {
				updatedLine += string(currLine[j])
			}
		}

			if lines[i] != updatedLine {
				lines[i] = updatedLine
				numLinesChanged++
			} 
		}

		if numLinesChanged == 0 {
			goAgain = false
		}		
	}
	println("Total valid rolls: ", totalRolls)
}

