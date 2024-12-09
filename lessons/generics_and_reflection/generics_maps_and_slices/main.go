package main

func process1[T []byte | [2]byte | string](value T) {
	_ = value[0]
}

func process2[T map[int]string](value T) {
	_ = value[0]
}

func process3[T map[int]string|[]byte](value T) {
	_ = value[0]
}