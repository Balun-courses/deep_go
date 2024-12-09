package main

func process1[T int32 | int64](value *T) {
	*value = *value + 1 // ok
}

func process2[T *int32 | *int64](value T) {
	*value = *value + 1 // compilation error
}

func process3[T *Int, Int int32 | int64](value T) {
	*value = *value + 1 // ok
}
