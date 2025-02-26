package main

import (
	"testing"
	"unicode/utf8"
)

func RunesToBytes(rs []rune) []byte {
	n := 0
	for _, r := range rs {
		n += utf8.RuneLen(r)
	}

	n, bs := 0, make([]byte, n)
	for _, r := range rs {
		n += utf8.EncodeRune(bs[n:], r)
	}

	return bs
}

// go test -bench=. -benchmem perf_test.go

var Result []byte
var str string = "€€€v€wqeqwwerrerrqw12313123€€€v€€€€v€€€€v€€€€v€"

func BenchmarkConversion(b *testing.B) {
	rs := []rune(str + "sdfsfs€€€v€€€€v€23423423")
	for i := 0; i < b.N; i++ {
		Result = RunesToBytes(rs)
	}
}

func BenchmarkConversionWithStr(b *testing.B) {
	rs := []rune(str + "sdfsfs€€€v€€€€v€23423423")
	for i := 0; i < b.N; i++ {
		Result = []byte(string(rs))
	}
}
