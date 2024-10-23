package main

func readFromNilMap() {
	var data map[int]int
	_ = data[100]
}

func deleteFromNilMap() {
	var data map[int]int
	delete(data, 100)
}

func writeToNilMap() {
	var data map[int]int
	data[100] = 100
}

func rangeByNilMap() {
	var data map[int]int
	for range data {
	}
}

func rewriteExistingKey() {
	data := make(map[int]int)
	data[100] = 500
	data[100] = 1000
}

func deleteNonExistingKey() {
	data := make(map[int]int)
	delete(data, 100)
}

func main() {
}
