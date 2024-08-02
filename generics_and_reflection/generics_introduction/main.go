package main

func GetKeys(data map[string]int) []string {
	keys := make([]string, 0, len(data))
	for key := range data {
		keys = append(keys, key)
	}

	return keys
}

func GetKeysGeneric[K comparable, V any](data map[K]V) []K {
	keys := make([]K, 0, len(data))
	for key := range data {
		keys = append(keys, key)
	}

	return keys
}

func main() {
	GetKeys(map[string]int{})
	GetKeysGeneric(map[int]int{})
	GetKeysGeneric[int, int](map[int]int{})
}
