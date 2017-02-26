package main

import "fmt"

func main() {

	printNPrimeNumbers(10)
	fmt.Println()
	findNthPrime(10)
}

func printNPrimeNumbers(n int) {
	count := 0
	i := 0
	for count < n {
		// for count < n {
		if isPrime(i) {
			fmt.Printf("%d is Prime\n", i)
			count++
		}
		i++
	}
}

func isPrime(number int) bool {
	if number < 2 {
		return false
	}
	if number == 2 {
		return true
	} else if number%2 == 0 {
		return false
	}
	for i := 3; i*i <= number; i = i + 2 {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func findNthPrime(n int) int {
	nthPrime := 0
	count := 0
	i := 0
	for count < n {
		if isPrime(i) {
			count++
			if count == n {
				fmt.Printf("%d is the %d Prime\n", i, n)
				nthPrime = i
				break
			}
		}
		i++
	}
	return nthPrime
}
