//PROBLEM STATEMENT
//2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

package main

import "log"

func main() {
	log.Println("starting problem 5")
	startingNumber := 1
	for {
		parseNum := getPureNumber(startingNumber)
		if parseNum != 0 {
			log.Println(parseNum)
			break
		}
		startingNumber++
		if startingNumber > 1000000000000000000 {
			break
		}
	}
}

func getPureNumber(number int) int {
	// var purity bool //every
	var status bool
	for i := 1; i < 21; i++ {
		if number%i == 0 {
			if number % 2 == 0 {
				status = true
			}
		} else {
			return 0
		}
		if status && i == 20 {
			return number
		}
	}

	return 0
}
