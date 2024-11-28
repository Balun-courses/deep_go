package main

func Process1[T any, K comparable, V any](key K, value V) {
	// implementation
}

func Process2[K comparable, V any, T any](key K, value V) {
	// implementation
}

func Process3[T any, _ comparable, _ any](value T) {
	// implementation
}

func Process4[T1, T2 any, _ comparable](value1 T1, value2 T2) {
	// implementation
}

func Process5[_ any, _ any]() {
	// implementation
}

func main() {
	Process1[int]("key", []int{100}) // partial argument list only with prefix
	Process2[int]("key", []int{100}) // compilation error with suffix

	Process3[int, string, float32](100)
	Process2[int, _, float32](100) // compilation error
}
