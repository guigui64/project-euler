package main

import "fmt"

func main() {
	fib := []int{1, 2}
	sum := 2
	for {
		fib = append(fib, fib[len(fib)-1]+fib[len(fib)-2])
		if fib[len(fib)-1] > 4000000 {
			break
		}
		if fib[len(fib)-1]%2 == 0 {
			sum += fib[len(fib)-1]
		}
	}
	fmt.Println(sum)
}
