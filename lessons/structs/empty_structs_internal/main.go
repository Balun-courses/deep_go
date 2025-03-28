package main

import "time"

func fn() {
	var v3 int
	var e3 struct{}

	println("v3:", &v3)
	println("e3:", &e3)
}

func main() {
	var v1 int
	var e1 struct{}
	var v2 int
	var e2 struct{}

	println("v1:", &v1)
	println("e1:", &e1)
	println("v2:", &v2)
	println("e2:", &e2)

	fn()

	time.Sleep(time.Second)

	go func() {
		var gv1 int
		var ge1 struct{}
		var gv2 int
		var ge2 struct{}

		println()
		println("gv1:", &gv1)
		println("ge1:", &ge1)
		println("gv2:", &gv2)
		println("ge2:", &ge2)
	}()

	time.Sleep(time.Second)
}
