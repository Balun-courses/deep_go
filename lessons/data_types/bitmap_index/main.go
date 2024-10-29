package main

// 0000 0001 -> есть кальяны
// 0000 0010 -> можно с животными
// 0000 0100 -> есть виранда
// 0000 1000 -> есть алкоголь
// 0001 0000 -> есть живая музыка

func searchRestaurants(pattern int8, bitmaps []int8) []int {
	var indexes []int
	for idx, bitmap := range bitmaps {
		if bitmap^pattern == 0 {
			indexes = append(indexes, idx)
		}
	}

	return indexes
}

func main() {
	restaurants := []int8{
		0b00001101,
		0b00000010,
		0b00010000,
		0b00011111,
		0b00001001,
	}

	pattern := int8(0b00011000)
	indexes := searchRestaurants(pattern, restaurants)
	_ = indexes
}
