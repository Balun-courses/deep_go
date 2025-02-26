package main

import (
	"testing"
	"unique"
)

// go test -bench=. bench_test.go

var str1 = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, 
	sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. 
	Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris 
	nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in 
	reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla 
	pariatur. Excepteur sint occaecat cupidatat non proident, sunt 
	in culpa qui officia deserunt mollit anim id est laborum`

var str2 = `Lorem ipsum dolor sit amet, consectetur adipiscing elit, 
	sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. 
	Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris 
	nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in 
	reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla 
	pariatur. Excepteur sint occaecat cupidatat non proident, sunt 
	in culpa qui officia deserunt mollit anim id est laborum`

var Result bool

// describe about:
// - space optimization
// - other types
// - synchronization

func BenchmarkConversion(b *testing.B) {
	handle1 := unique.Make(str1 + "!!!")
	handle2 := unique.Make(str2 + "!!!")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Result = handle1 == handle2
	}
}

func BenchmarkConversionWithStr(b *testing.B) {
	sstr1 := str1 + "!!!"
	sstr2 := str2 + "!!!"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Result = sstr1 == sstr2
	}
}
