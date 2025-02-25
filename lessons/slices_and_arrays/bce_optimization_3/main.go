package main

// go test -gcflags="-d=ssa/check_bce" -bench=. bench_test.go

func sum1(values []int, size int) int {
	sum := 0
	for i := 0; i < size; i++ {
		sum += values[i]
	}

	return sum
}

func sum2(values []int, size int) int {
	_ = values[size-1] // for BCE

	sum := 0
	for i := 0; i < size-1; i++ {
		sum += values[i]
	}

	return sum
}

func main() {

}
