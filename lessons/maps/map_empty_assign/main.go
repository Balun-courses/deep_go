package main

func main() {
	m := make(map[string]int, 3)
	x := len(m)

	m["Go"] = m["Go"]

	y := len(m)
	println(x, y)
}
