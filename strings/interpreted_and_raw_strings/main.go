package main

import "fmt"

func main() {
	interpreted := "\nHello\nworld!\n"
	raw := `\nHello\nworld!\n`

	fmt.Println(interpreted)
	fmt.Println(raw)

	interpretedData := []byte("\n")
	rawData := []byte(`\n`)

	fmt.Println("interpretedData:", interpretedData)
	fmt.Println("rawData:", rawData)
}
