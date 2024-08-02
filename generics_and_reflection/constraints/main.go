package main

import "strconv"

// Need to show solution

type MyConstraint interface {
	int
	String() string
}

type MyInt int

func (i MyInt) String() string {
	return strconv.Itoa(int(i))
}

func GetKeysGeneric[K MyConstraint, V any](data map[K]V) []K {
	keys := make([]K, 0, len(data))
	for key := range data {
		keys = append(keys, key)
	}

	return keys
}

func main() {
	GetKeysGeneric(map[MyInt]int{})
}
