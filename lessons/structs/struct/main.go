package main

type Data struct {
	x int
	y int
}

func main() {
	// 1 way
	data1 := Data{10, 20}
	_ = data1

	// 2 way
	data2 := Data{x: 10, y: 20}
	_ = data2
}
