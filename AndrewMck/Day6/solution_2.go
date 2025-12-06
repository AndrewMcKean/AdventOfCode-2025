package main

import (
	"bufio"
	"strings"
	"strconv"
	"os"
	"slices"
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

	operatorLine := lines[len(lines) - 1]
	splitPoints := []int{}

	for i := 0; i < len(operatorLine); i++ {
		char := string(operatorLine[i])

		if char == "*" || char == "+" {
			if i > 0 {
				splitPoints = append(splitPoints, i - 1)
			}
		}
	}

	numArrays := [][]string{}
	numArray := []string{}

	operand := "" 

	for col := 0; col < len(operatorLine); col++ {
		if slices.Contains(splitPoints, col) {
			numArrays = append(numArrays, slices.Clone(append(numArray, operand)))
			numArray = numArray[:0]
			continue
		}

		colStr := ""

		for _, line := range lines {
			char := string(line[col])
			if char == "*" || char == "+" {
				operand = char 
			} else {
				colStr += char 
			}
		}
		numArray = append(numArray, strings.Trim(colStr, " "))

		if col == len(operatorLine) - 1 {
			numArrays = append(numArrays, slices.Clone(append(numArray, operand)))
		}
	}

	sum := 0
	for _, numArray := range numArrays {

		problemElements := []string{}
		operation := numArray[len(numArray)- 1]	

		for _, num := range numArray {
			problemElements = append(problemElements, num) 
		}

		if operation == "+" {
			sum += addNums(problemElements)
		} else {
			sum += multiplyNums(problemElements)
		}
	}

	println("Total: ", sum)
}
