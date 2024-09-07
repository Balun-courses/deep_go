package main

import (
	"fmt"
	"hash/maphash"
)

func main() {
	var h1 maphash.Hash
	h1.WriteString("hello")
	fmt.Println(h1.Sum64(), h1.Seed()) // random hash seed

	var h2 maphash.Hash
	h2.WriteString("hello")
	fmt.Println(h2.Sum64(), h2.Seed()) // random hash seed

	var h3 maphash.Hash
	h3.SetSeed(h2.Seed())
	h3.WriteString("hello")
	fmt.Println(h3.Sum64(), h3.Seed()) // hash seed from h2
}
