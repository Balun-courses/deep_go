package main

import "math"

func IncInt(counter int) int {
	if counter == math.MaxInt {
		panic("int overflow")
	}

	return counter + 1
}

func IncUint(counter uint) uint {
	if counter == math.MaxUint {
		panic("uint overflow")
	}

	return counter + 1
}

func AddInt(lhs, rhs int) int {
	if lhs > math.MaxInt-rhs {
		panic("int overflow")
	}

	return lhs + rhs
}

func MulInt(lhs, rhs int) int {
	if lhs == 0 || rhs == 0 {
		return 0
	}

	if lhs == 1 || rhs == 1 {
		return lhs * rhs
	}

	if lhs > math.MaxInt/rhs {
		panic("int overflow")
	}

	/*
		int a = <something>;
		int x = <something>;
		// There may be a need to check for -1 for two's complement machines.
		// If one number is -1 and another is INT_MIN, multiplying them we get abs(INT_MIN) which is 1 higher than INT_MAX
		if (a == -1 && x == INT_MIN) // `a * x` can overflow
		if (x == -1 && a == INT_MIN) // `a * x` (or `a / x`) can overflow
		// general case
		if (x != 0 && a > INT_MAX / x) // `a * x` would overflow
		if (x != 0 && a < INT_MIN / x) // `a * x` would underflow
	*/

	return lhs * rhs
}
