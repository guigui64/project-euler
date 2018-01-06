package main

import "fmt"

func isPrime(n int) bool {
	for d := 2; d < n; d++ {
		if n%d == 0 {
			return false
		}
	}
	return true
}

func maxPrimeFactor(n int) int {
	if isPrime(n) {
		return n
	}
	for f := n / 2; f > 1; f-- {
		if n%f == 0 && isPrime(f) {
			return f
		}
	}
	return -1
}

func main() {
	fmt.Println(maxPrimeFactor(13195))
	fmt.Println(maxPrimeFactor(600851475143))
}
