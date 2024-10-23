package main

func main() {
	data := map[interface{}]int{
		"a": 1,
		2:   2,
	}

	data[[]int{}] = 3
	data[func() {}] = 4
}
