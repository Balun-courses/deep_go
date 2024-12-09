package main

func process[T ~string | ~[]byte](value T, index int) {
	_ = value[index]
	value[index] = byte('a') // error
}
