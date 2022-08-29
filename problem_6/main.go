//PROBLEM STATEMENT
// The sum of the squares of the first ten natural numbers is,

// The square of the sum of the first ten natural numbers is,

// Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is .

// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

package main

import "log"

func main() {
	log.Println("starting prtoblem six")
	sumOfSquareOfNaturalNumber := sumOfSquaresOfNaturalNumbers(100)
	sumOfNaturalNumberSquare := sumOfNaturalNumbersSquare(100)

	res := sumOfNaturalNumberSquare - sumOfSquareOfNaturalNumber

	log.Println(res)
}

func sumOfSquaresOfNaturalNumbers(number int) int {
	var square, sum int
	for i := 0; i <= number; i++ {
		square = i * i
		sum += square
	}
	return sum
}

func sumOfNaturalNumbersSquare(number int) int {
	var square, sum int
	for i := 0; i <= number; i++ {
		sum += i
	}
	square = sum * sum

	return square
}
