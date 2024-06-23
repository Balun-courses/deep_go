package main

import "fmt"

var table = map[uint8]byte{
	97: 'a', 98: 'b', 99: 'c', 100: 'd', 101: 'e',
	102: 'f', 103: 'g', 104: 'h', 105: 'i', 106: 'j',
	107: 'k', 108: 'l', 109: 'm', 110: 'n', 111: 'o',
	112: 'p', 113: 'q', 114: 'r', 115: 's', 116: 't',
	117: 'u', 118: 'v', 119: 'w', 120: 'x', 121: 'y',
	122: 'z',
}

func printText(text []uint8) {
	for _, codeNumber := range text {
		symbol, found := table[codeNumber]
		if !found {
			fmt.Printf("#incorrect code symbol: %d\n", codeNumber)
			return
		}

		fmt.Printf("%q", symbol)
	}
}

func main() {
	text := []uint8{104, 101, 108, 108, 111}
	printText(text)
}
