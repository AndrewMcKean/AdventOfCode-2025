package main

import (
"bufio"
"os"
"strings"
"strconv"
)

func getFactors(num int) []int {
	factors := []int{}

	for i := 1; i < num; i++ {
		if num % i == 0 {
			factors = append(factors, i)	
		}
	}

	return factors
}

func main () {
	fi, _ := os.Open("inputTbox.txt")
	scanner := bufio.NewScanner(fi)

	scanner.Scan()
	
	var csv = scanner.Text()
	
	var IDs = strings.Split(csv, ",")

	var count = 0

	for _, v := range IDs {
		var parts = strings.Split(v, "-")
		var start, _ = strconv.Atoi(parts[0])
		var end, _ = strconv.Atoi(parts[1])

		for i := start; i < end + 1; i++ {
			var numberToAssess = strconv.Itoa(i)

			var factors = getFactors(len(numberToAssess))

			for factorIndex := 0; factorIndex < len(factors); factorIndex++ {
				var factor = factors[factorIndex]

				var invalidID = true

				var totalChunks = len(numberToAssess) / factor
				
				var prevChunk = numberToAssess[:factor]
				for chunk := 1; chunk < totalChunks; chunk++ {
					var strChunkToAssess = numberToAssess[factor * chunk:factor * (chunk + 1)]
					if prevChunk != strChunkToAssess {
						invalidID = false	
						break
					} 
				}

				if invalidID == true {
					count += i
					break
				}
			}
		}
	}

	println("Count: ", count)
}

