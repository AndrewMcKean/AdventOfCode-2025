package main

import (
	"bufio"
	"os"
	"strconv"
)

func scaryRecursiveFunction(string searchString int currentIndex string stringSoFar) {

	if len(stringSoFar) == 12 {
		// Done
		return stringSoFar
	}

	// Seach logic
	currentBiggest := 0
	currentBiggestIndex := 0

	for i := 0; i < len(searchString) - currentIndex; i++ {
		charAsNum, _ := strconv.Atoi(string(searchString[i]))
		
		if charAsNum > currentBiggest {
			currentBiggestIndex = i
			currentBiggest = charAsNum
		}
	}

	// Prepare recursive call
	nextString := searchString[currentBiggestIndex:currentIndex-1]
	newStringSoFar := stringSoFar + strconv.Itoa(currengBiggest)
	newIndex := currentIndex - 1

	return scaryRecursiveFunction(nextString, newIndex, newStringSoFar) 
}

func main() {
	fi, _ := os.Open("inputTbox.txt")

  scanner := bufio.NewScanner(fi)
	
	total := 0

	const BATTERY_SIZE = 12

  for scanner.Scan() {
		line := scanner.Text()

		currentBiggest := 0
		currentBiggestIndex := 0
		
		secondBiggest := 0

		for i := 0; i < len(line) - 1; i++ {
			charAsNum, _ := strconv.Atoi(string(line[i]))

			if charAsNum > currentBiggest {
				currentBiggest = charAsNum 
				currentBiggestIndex = i	
			}
		}

		for i := currentBiggestIndex; i < len(line); i++ {
			charAsNum, _ := strconv.Atoi(string(line[i]))

			if i != currentBiggestIndex && charAsNum > secondBiggest {
				secondBiggest = charAsNum
			}
		}

		totalForLine, _ := strconv.Atoi(strconv.Itoa(currentBiggest) + strconv.Itoa(secondBiggest))

		println(totalForLine)

		total += totalForLine

		println("Biggest in line: ", currentBiggest, secondBiggest)

	}	
	

	println(total)
}
