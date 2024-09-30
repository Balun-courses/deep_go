package main

import "testing"

// go test -bench=. comparison_test.go

type account struct {
	balance int
}

func BenchmarkWithPointers(b *testing.B) {
	accounts := [...]*account{
		{balance: 100},
		{balance: 200},
		{balance: 300},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, a := range accounts {
			a.balance += 1
		}
	}
}

func BenchmarkWithIndices(b *testing.B) {
	accounts := [...]account{
		{balance: 100},
		{balance: 200},
		{balance: 300},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := range accounts {
			accounts[i].balance += 1
		}
	}
}
