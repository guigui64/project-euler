package main

import "fmt"

func isPalindrome(n int) bool {
	s := fmt.Sprint(n)
	for shift := 0; shift < len(s)/2; shift++ {
		if s[shift] != s[len(s)-shift-1] {
			return false
		}
	}
	return true
}

func main() {
	bound := 999
	max := 0
	for a := bound; a > 1; a-- {
		for b := bound; b > 1; b-- {
			if isPalindrome(a*b) && a*b > max {
				max = a * b
			}
		}
	}
	fmt.Println(max)
}
