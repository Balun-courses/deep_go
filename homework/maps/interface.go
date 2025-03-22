package main

import "cmp"

type TreeMap[K cmp.Ordered, V any] interface {
	Insert(k K, v V)
	Erase(k K)
	Contains(k K) bool
	Size() int
	ForEach(action func(k K, v V))
}
