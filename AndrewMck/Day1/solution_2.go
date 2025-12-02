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
    
		fi, _ := os.Open("./inputTbox.txt") 

		scanner := bufio.NewScanner(fi)

		for scanner.Scan() {
			line := scanner.Text()

			direction := string(line[0])
			distance, _ := strconv.Atoi(string(line[1:]))

			if direction == "L" {
				for range distance {
					if current == 0 {
						current = 99
					} else {
            if current - 1 == 0 {
							countOfZeroes++
						}
						current--
					}
				}
			}

			if direction == "R" {
				for range distance {
					if current == 99 {
						current = 0
						countOfZeroes++
					} else {
						current++
					}
				} 
			}
		}

		println("Password: ", countOfZeroes)

}
