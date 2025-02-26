package main

import "unsafe"

func main() {
	const str1 string = "example"
	const str2 string = "example"
	const str3 string = "another example"

	println(unsafe.StringData(str1))
	println(unsafe.StringData(str2))
	println(unsafe.StringData(str3))
}
