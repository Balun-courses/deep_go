package main

import (
	"fmt"
	"math/rand"
	"testing"
)

// go test -bench=. comparison_test.go

type OODStyle struct {
	Field1 int
	Field2 string
	Field3 int
	Field4 string
	Field5 int
	Field6 string
	Field7 int
	Field8 string
}

type DODStyle struct {
	Field1 []int
	Field2 []string
	Field3 []int
	Field4 []string
	Field5 []int
	Field6 []string
	Field7 []int
	Field8 []string
}

var Sink int

func BenchmarkDOD(b *testing.B) {
	r := rand.New(rand.NewSource(42))
	data := generateDOD(r, 1_000_000)
	b.ResetTimer()

	Sink = 0
	for i := 0; i < b.N; i++ {
		for j, f1 := range data.Field1 {
			if f1 == 500000 {
				Sink = j
			}
		}
	}
}

func BenchmarkOOD(b *testing.B) {
	r := rand.New(rand.NewSource(42))
	data := generateOOD(r, 1_000_000)
	b.ResetTimer()

	Sink = 0
	for i := 0; i < b.N; i++ {
		for j, ant := range data {
			if ant.Field1 == 500000 {
				Sink = j
			}
		}
	}
}

func generateOOD(r *rand.Rand, size int) []OODStyle {
	data := make([]OODStyle, size)
	for i := range data {
		data[i] = OODStyle{
			Field1: r.Intn(1000),
			Field2: fmt.Sprintf("field2-%d", r.Intn(1000)),
			Field3: r.Intn(1000),
			Field4: fmt.Sprintf("field4-%d", r.Intn(1000)),
			Field5: r.Intn(1000),
			Field6: fmt.Sprintf("field6-%d", r.Intn(1000)),
			Field7: r.Intn(1000),
			Field8: fmt.Sprintf("fiel8-%d", r.Intn(1000)),
		}
	}

	return data
}

func generateDOD(r *rand.Rand, size int) DODStyle {
	data := DODStyle{
		Field1: make([]int, size),
		Field2: make([]string, size),
		Field3: make([]int, size),
		Field4: make([]string, size),
		Field5: make([]int, size),
		Field6: make([]string, size),
		Field7: make([]int, size),
		Field8: make([]string, size),
	}

	for i := 0; i < size; i++ {
		data.Field1[i] = r.Intn(1000)
		data.Field2[i] = fmt.Sprintf("field2-%d", r.Intn(1000))
		data.Field3[i] = r.Intn(1000)
		data.Field4[i] = fmt.Sprintf("field4-%d", r.Intn(1000))
		data.Field5[i] = r.Intn(1000)
		data.Field6[i] = fmt.Sprintf("field6-%d", r.Intn(1000))
		data.Field7[i] = r.Intn(1000)
		data.Field8[i] = fmt.Sprintf("fiel8-%d", r.Intn(1000))
	}

	return data
}
