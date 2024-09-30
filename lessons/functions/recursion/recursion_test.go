package main

import "testing"

func SlowFactorial(number int) int {
	if number == 0 {
		return 1
	}

	return number * SlowFactorial(number-1)
}

func FastFactorial(number int) int {
	return factorial(number, 1)
}

func factorial(number, accumulator int) int {
	if number == 1 {
		return accumulator
	}

	return factorial(number-1, number*accumulator)
}

func BenchmarkSlowRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = SlowFactorial(10)
	}
}

func BenchmarkFastRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = FastFactorial(10)
	}
}
