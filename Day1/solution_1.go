package main

import (
	"bufio"
	"os"
	"strconv"
)

// We don't like to say "I'm dumb" - but I'm dumb. Stupid maths problems
// If time, do a bit of cleanup - compute the second condition first - then check 
// that against 0 or 100

func main () {
	var current = 50
	var countOfZeroes = 0

	fi, _ := os.Open("./input.txt")

	scanner := bufio.NewScanner(fi)

  for scanner.Scan() {
		line := scanner.Text()

		direction := string(line[0])
		distance, _ := strconv.Atoi(string(line[1:]))

		println("Current: ", current)
		println(line)

		if direction == "L" {
			// Minus operation
			remainder := (current - distance) % 100
			endsAtZero := current - distance == 0 || remainder == 0

			if endsAtZero {
				countOfZeroes++
			}

			current = remainder

			if remainder < 0 {
				current += 100
			}
		}

		if direction == "R" {
			// Plus operation
			remainder := (current + distance) % 100

			if remainder == 0 {
				countOfZeroes++
			}

			current = remainder
		}

		println("Running countOfZeroes: ", countOfZeroes)
	}

	println("Password: ", countOfZeroes)
}
