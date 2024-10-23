package main

func main() {
	data1 := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	data2 := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}

	for key := range data1 {
		print(key)
	}

	println()

	for key := range data2 {
		print(key)
	}

	println()
}
