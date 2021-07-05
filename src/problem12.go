package main

import (
	"fmt"
	"math"
)

func problem12() {
	calculateTriangleNumberWithNDivisors(500)
}

func calculateTriangleNumberWithNDivisors(n int) {
	size := 0
	ittr := 1
	triangleNumber := 0

	for size <= n {
		triangleNumber = calculateTriangleNumber(ittr)
		divisors := calculateDivisors(float64(triangleNumber))
		size = len(divisors)
		if ittr%1000 == 0 {
			fmt.Println(ittr, size)
		}
		ittr++
	}

	fmt.Println("Triangle number: ", triangleNumber)
	fmt.Println("Num Elements: ", size)
	fmt.Println("Iterations taken: ", ittr-1)
}

func calculateTriangleNumber(ittr int) int {
	var x = 0

	for i := 1; i <= ittr; i++ {
		x += i
	}

	return x
}

func calculateDivisors(triNum float64) []float64 {
	var divisors []float64

	var i float64 = 2

	divisors = append(divisors, 1, triNum)

	for i <= math.Sqrt(triNum) {
		if math.Mod(triNum, i) == 0 {
			divisors = append(divisors, i)
			if i != (triNum / i) {
				divisors = append(divisors, triNum/i)
			}
		}
		i++
	}

	return divisors
}
