package main

func Fibonacci(number int) int {
	if number <= 2 {
		return 1
	}

	return Fibonacci(number-1) + Fibonacci(number-2)
}

func FibonacciWithMemoization(number int) int {
	cache := make([]int, number+1)
	var impl func(number int) int
	impl = func(n int) int {
		if cache[n] != 0 {
			return cache[n]
		}

		if number <= 2 {
			return 1
		} else {
			cache[n] = impl(n-1) + impl(n-2)
		}

		return cache[n]
	}

	return impl(number)
}
