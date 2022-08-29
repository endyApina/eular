//PROBLEM STATEMENT
//The prime factors of 13195 are 5, 7, 13 and 29.

// What is the largest prime factor of the number 600851475143 ?

package main

import (
	"log"
	"math"
	"sort"
)

func main() {
	log.Println("starting")
	number := 600851475143
	primeFactors := getPrimeFactors(number)
	log.Println(primeFactors)
	highestPrimeFactor := getHighestNumberInArray(primeFactors)
	log.Println(highestPrimeFactor)
}

func getPrimeFactors(number int) []int {
	primeFactors := []int{}

	for number%2 == 0 {
		primeFactors = append(primeFactors, 2)
		number = number / 2
	}

	for i := 3; i <= int(math.Sqrt(float64(number))); i = i + 2 {
		for number%i == 0 {
			primeFactors = append(primeFactors, i)
			number = number / i
		}
	}

	if number > 2 {
		primeFactors = append(primeFactors, number)
	}
	return primeFactors
}

func getHighestNumberInArray(arrayObject []int) int {
	sort.Ints(arrayObject)
	return arrayObject[len(arrayObject)-1]
}
