package main

import (
	"fmt"
	"testing"
)

func rangeWithoutAllocation() {
	var str = "world"
	for range []byte(str) { // no allocation with copy
	}
}

func workWithMapsWithoutAllocation() {
	key := []byte{'k', 'e', 'y'}

	data := map[string]string{}

	data[string(key)] = "value" // allocation with copy
	_ = data[string(key)]       // no allocation with copy
}

func comparisonAndConcatenation1() {
	var x = []byte{1023: 'x'}
	var y = []byte{1023: 'y'}

	if string(x) != string(y) { // no allocation with copy
		s := (" " + string(x) + string(y))[1:] // no allocation with copy
		_ = s
	}
}

func comparisonAndConcatenation2() {
	var x = []byte{1023: 'x'}
	var y = []byte{1023: 'y'}

	if string(x) != string(y) { // no allocation with copy
		s := string(x) + string(y) // allocation with copy
		_ = s
	}
}

func main() {
	fmt.Println(testing.AllocsPerRun(1, comparisonAndConcatenation1)) // 1
	fmt.Println(testing.AllocsPerRun(1, comparisonAndConcatenation2)) // 3
}
