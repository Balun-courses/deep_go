package main

func main() {
	data := map[int]int{100: 10000}
	value := data[100]
	_ = value

	// under hood
	// pk := unsafe.Pointer(&key)
	// pv := runtime.mapaccess1(typeOf(data), data, pk)
	// value := *(*int)(pv)
}
