package main

import (
	"bufio"
	"strings"
  "strconv"
	"os"
)

func addNums(problemElements []string) int {
	sum := 0
	for _, num := range problemElements {
		number, err := strconv.Atoi(num)

		if err != nil {
			continue
		}
		
		sum += number
	}

	return sum
}

func multiplyNums(problemElements []string) int {
	product := 1
	for _, num := range problemElements {
		number, err := strconv.Atoi(num)

		if err != nil {
			continue
		}

		product *= number
	}
	return product
}

func main () {
	fi, _ := os.Open("inputTbox.txt")
	scanner := bufio.NewScanner(fi)
	
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	lineArrays := [][]string{}
	for _, line := range lines {
		stringsSplit := strings.Split(line, " ")

		numArray := []string{}
		for _, element := range stringsSplit {
			trimmedElement := strings.Trim(element, " ")
			if trimmedElement != "" {
				numArray = append(numArray, trimmedElement)
				println(trimmedElement)
			}
		}
	
		lineArrays = append(lineArrays, numArray)	
	}

	operations := lineArrays[len(lineArrays) - 1]

	sum := 0
	for i, operation := range operations {
	
		problemElements := []string{}
		
		for _, lineArray := range lineArrays {
			problemElements = append(problemElements, lineArray[i])
		}
		
		if operation == "+" {
			sum += addNums(problemElements)
		} else {
			sum += multiplyNums(problemElements)
		}
	}

	println("Total: ", sum)

}
