package main

import (
"bufio"
"os"
"strings"
"strconv"
)

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

			println(i, numberToAssess)

			if len(numberToAssess) % 2 == 0 {
				var splitPoint = len(numberToAssess) / 2

				var sequenceA = numberToAssess[:splitPoint]
				var sequenceB = numberToAssess[splitPoint:]

				println("A: ", sequenceA, "B: ", sequenceB)

				if sequenceA == sequenceB {
					count += i 
				}
			}
		}
	}

	println("Count: ", count)
} 
