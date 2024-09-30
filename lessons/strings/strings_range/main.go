package main

import "fmt"

func main() {
	text := "Sr, привет 世界"
	for idx, symbol := range text { // range []rune(text)
		fmt.Printf("%d-%c ", idx, symbol)
	}

	fmt.Println()
	for i := 0; i < len(text); i++ { // range []byte(text)
		fmt.Printf("%d-%c ", i, text[i])
	}
}
