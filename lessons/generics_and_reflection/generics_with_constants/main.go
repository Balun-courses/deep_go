package main

func process1[T int]() {
	const value T = 5 // compilation error
}

func process2[T int]() {
	var value T = 5 // ok
	_ = value
}
