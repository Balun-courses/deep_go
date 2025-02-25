package main

// go run -gcflags="-d=ssa/check_bce" main.go

func fn1(data []int, check func(int) bool) []int {
	var idx = 0
	for _, value := range data {
		if check(value) {
			data[idx] = value // Found IsInBounds
			idx++
		}
	}

	return data[:idx] // Found IsSliceInBounds
}

func fn2(lhs, rhs []byte) {
	for idx := range min(len(lhs), len(rhs)) {
		_ = lhs[idx] // Found IsInBounds
		_ = rhs[idx] // Found IsInBounds
	}
}

func fn3(data [256]byte) {
	for idx := 0; idx < 128; idx++ {
		_ = data[2*idx] // Found IsInBounds
	}
}

func main() {}
