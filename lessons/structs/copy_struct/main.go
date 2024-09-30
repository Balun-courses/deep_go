package main

type Data struct {
	x int
	y int
}

func main() {
	// 1 way
	data1 := Data{x: 10, y: 10}
	temp1 := data1
	_ = temp1

	// 2 way
	data2 := Data{x: 10, y: 10}
	var temp2 Data
	temp2.x = data2.x
	temp2.y = data2.y
}
