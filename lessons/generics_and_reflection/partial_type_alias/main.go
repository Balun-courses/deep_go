package main

type T[T1 any, T2 any] struct{}

type TypeAlias1 = T[string, int]
type TypeAlias2[T1 any] = T[T1, int]

type TypeDefinition1 T[string, int]
type TypeDefinition2[T1 any] T[T1, int]
