package main

import "log"

func factorial(n int) int {
	res := 1

	for i := 1; i <= n; i++ {
		res *= i
	}

	return res
}

func countVowelStrings(n int) int {
	vowelsFact := factorial(5)

	return vowelsFact / (factorial(5-n) * factorial(n))
}

//n = 5

//m!
//-
//(m−n)!⋅n!

func main() {
	obj := countVowelStrings(2)

	log.Println(obj)
}
