package main

import (
	"os"
	"bufio"
	"strings"
	"slices"
)

func appendAsSet(slice []int, target int) []int {
	if slices.Contains(slice, target) {
		return slice
	} else {
		return append(slice, target)
	}	
}

func main() {
	fi, _ := os.Open("inputArch.txt")
	scanner := bufio.NewScanner(fi)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	existingBeams := []int{strings.Index(lines[0], string("S"))}
	splits := 0

	// Loose algo
	// Calc num unique roots for EACH node
	// For any given node, the number of unique routes is equal to 
	// the number of unique routes of it's parents + 2

	for lineIndex, s := range lines {
		splitIndicesForLine := []int{}
		for charIndex, char := range s {
			if string(char) == "^" {
				if slices.Contains(existingBeams, charIndex) {
					// Split it - i.e remove it, and create a beam either side
					existingBeamIndex := slices.Index(existingBeams, charIndex)
					existingBeams = slices.Delete(existingBeams, existingBeamIndex, existingBeamIndex + 1)
					
					existingBeams = appendAsSet(existingBeams, charIndex - 1)
					existingBeams = appendAsSet(existingBeams, charIndex + 1)

					splitIndicesForLine = appendAsSet(splitIndicesForLine, charIndex - 1)
					splitIndicesForLine = appendAsSet(splitIndicesForLine, charIndex + 1)
				}
			}
		}

		println(splits, " splits for line: ", lineIndex)
		// splits += len(splitIndicesForLine)

	}

	println("Total splits: ", splits)
}
