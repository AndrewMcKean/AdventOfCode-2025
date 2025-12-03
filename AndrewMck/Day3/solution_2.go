package main

import (
	"bufio"
	"os"
	"strconv"
)

const BATTERY_SIZE = 12

func scaryRecursiveFunction(searchString string, currentIndex int, stringSoFar string) string {

	println("StrSoFar: " , stringSoFar)
	println("SearchString: ", searchString)
	if len(stringSoFar) == BATTERY_SIZE {
		// Done
		return stringSoFar
	}

	// Seach logic
	currentBiggest := 0
	currentBiggestIndex := 0

	lastPossibleIndex := len(searchString) - currentIndex + 1

	println("LastPossibleIndex: ", lastPossibleIndex, " for currentIndex: ", currentIndex)

	for i := 0; i < lastPossibleIndex; i++ {
		charAsNum, _ := strconv.Atoi(string(searchString[i]))
	
		println("Comparing: ", charAsNum, " vs. ", currentBiggest)
		if charAsNum > currentBiggest {
			currentBiggestIndex = i
			currentBiggest = charAsNum
		}
	}

	// Prepare recursive call
	nextString := searchString[currentBiggestIndex + 1:]
	newStringSoFar := stringSoFar + strconv.Itoa(currentBiggest)
	newIndex := currentIndex - 1

	return scaryRecursiveFunction(nextString, newIndex, newStringSoFar) 
}

func main() {
	fi, _ := os.Open("inputTbox.txt")

  scanner := bufio.NewScanner(fi)
	
	total := 0

  for scanner.Scan() {
		line := scanner.Text()
		
		totalForLine, _ := strconv.Atoi(scaryRecursiveFunction(line, BATTERY_SIZE, ""))

		total += totalForLine

		println("TotalForLine: ", totalForLine)
	}	
	

	println("Total: ", total)
}
