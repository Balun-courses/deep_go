package main

import "fmt"

func main() {
	var str = "world"
	// Here, the []byte(str) conversion will
	// not copy the underlying bytes of str.
	for i, b := range []byte(str) {
		fmt.Println(i, ":", b)
	}

	key := []byte{'k', 'e', 'y'}
	m := map[string]string{}
	// The string(key) conversion copys the bytes in key.
	m[string(key)] = "value"
	// Here, this string(key) conversion doesn't copy
	// the bytes in key. The optimization will be still
	// made, even if key is a package-level variable.
	fmt.Println(m[string(key)]) // value (very possible)
}

/*
var s string
var x = []byte{1023: 'x'}
var y = []byte{1023: 'y'}

func fc() {
	// None of the below 4 conversions will
	// copy the underlying bytes of x and y.
	// Surely, the underlying bytes of x and y will
	// be still copied in the string concatenation.
	if string(x) != string(y) {
		s = (" " + string(x) + string(y))[1:]
	}
}

func fd() {
	// Only the two conversions in the comparison
	// will not copy the underlying bytes of x and y.
	if string(x) != string(y) {
		// Please note the difference between the
		// following concatenation and the one in fc.
		s = string(x) + string(y)
	}
}

func main() {
	fmt.Println(testing.AllocsPerRun(1, fc)) // 1
	fmt.Println(testing.AllocsPerRun(1, fd)) // 3
}
*/
