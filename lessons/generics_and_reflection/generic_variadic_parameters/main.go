package main

func Process[T any](data ...T) {
	// processing
}

func main() {
	Process(2, 5, 10)
	Process("one", "two")
	Process(0.3)
}
