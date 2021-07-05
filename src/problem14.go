package main

import "fmt"

var recursions = 1

func problem14() {
	value := 0
	best := 0
	for i := 1; i < 1000000; i++ {
		collatz(i)
		if recursions > best {
			best = recursions
			value = i
		}
		recursions = 1
	}

	fmt.Println(best)
	fmt.Println(value)
}

func collatz(n int) int {
	recursions++

	if n%2 == 0 {
		n /= 2
	} else {
		n = 3*n + 1
	}

	if n == 1 {
		return 1
	} else {
		return collatz(n)
	}
}
