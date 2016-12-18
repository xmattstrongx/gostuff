package main

import "fmt"

type primeChecker struct {
	size        int
	primeChecks []int
}

func main() {
	checks := primeChecker{}
	checks.readInData()
	checks.printPrimeEvaluationResults()
}

func (p *primeChecker) readInData() {
	fmt.Scanf("%d", &p.size)

	p.primeChecks = make([]int, p.size)
	for i := range p.primeChecks {
		fmt.Scanf("%d", &p.primeChecks[i])
	}
}

func (p *primeChecker) printPrimeEvaluationResults() {
	for _, num := range p.primeChecks {
		if isPrime(num) {
			fmt.Println("Prime")
		} else {
			fmt.Println("Not prime")
		}
	}
}

// If a number is divisable by another number
// less or equal to the square root of the first number...
// it is NOT prime. That is the reason for O(sqrt(n)) run time.
func isPrime(n int) bool {
	if n == 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
