//PROBLEM STATEMENT
//A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

// Find the largest palindrome made from the product of two 3-digit numbers.

package main

import (
	"log"
	"sort"
)

type productsObject struct {
	products          []int
	palindromicNumber int
}

func main() {
	log.Println("Hello here")
	palindromicNumberArray := []productsObject{}
	palindNum := []int{}
	for i := 100; i < 1000; i++ {
		for y := 100; y < 1000; y++ {
			product := i * y
			isPalindromic := isPalindromicNumber(product)
			if isPalindromic {
				productsArray := []int{i, y}
				palindromicResponseObject := productsObject{
					products:          productsArray,
					palindromicNumber: product,
				}
				palindNum = append(palindNum, product)
				palindromicNumberArray = append(palindromicNumberArray, palindromicResponseObject)
			}
		}

	}
	sort.Ints(palindNum)
	log.Println(palindromicNumberArray[len(palindromicNumberArray)-1])
	log.Println(palindNum[len(palindromicNumberArray)-1])

}

func isPalindromicNumber(number int) bool {
	inputNumber := number
	remainder := 0
	response := 0

	for number > 0 {
		remainder = number % 10
		response = (response * 10) + remainder
		number = number / 10
	}

	if inputNumber == response {
		return true
	} else {
		return false
	}
}
