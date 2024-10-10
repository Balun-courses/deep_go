package main

// -l = disable inlining
// -m = print optimization decisions
// go build -gcflags '-l -m'
// go run -gcflags '-l' main.go

func process(index int) byte {
	var data [1 << 20]byte
	return data[index]
}

func main() {
	var index int = 100

	pointer := &index
	println("pointer:", pointer)
	process(index)
	println("pointer:", pointer)
}
