package main

import "unsafe"

func action() {}

func main() {
	var str = "go"
	newStr := str + "-go"

	strData := unsafe.StringData(str)
	newStrData := unsafe.StringData(newStr)

	println("action:", action)
	println("strData:", strData)
	println("newStrData:", newStrData)

	slice := unsafe.Slice(strData, len(str))
	newSlice := unsafe.Slice(newStrData, len(newStr))

	newSlice[0] = 'G'
	println("newStr:", newStr)

	slice[0] = 'G'
	println("str:", str)
}
