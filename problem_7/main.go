package main

import (
	"log"
)

func main() {
	log.Println("starting problem 7")
	startingNumber := 1
	primeNumbers := []int{1}
	var response int
	for {
		isAprimeNumber := isPrime(startingNumber)
		if isAprimeNumber {
			primeNumbers = append(primeNumbers, startingNumber)
		}

		startingNumber++
		lenOfPrimeNumbers := len(primeNumbers)
		if lenOfPrimeNumbers > 10001 {
			response = primeNumbers[lenOfPrimeNumbers-1]
			break
		}
	}
	log.Println(primeNumbers)
	log.Println(response)
}

func isPrime(number int) bool {
	for i := 2; i <= number; i++ {
		if number%i == 0 {
			if i == number {
				return true
			}
			return false
		}
	}
	return false
}
