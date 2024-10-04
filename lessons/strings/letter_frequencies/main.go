package main

import "fmt"

func CalculateFrequencies(str string) {
	frequencies := [26]int{}
	for _, letter := range str {
		letterIdx := letter - 'a'
		frequencies[letterIdx]++
	}

	for letterIdx, frequency := range frequencies {
		if frequency != 0 {
			letter := 'a' + letterIdx
			fmt.Printf("%c = %d\n", letter, frequency)
		}
	}
}

func main() {
	CalculateFrequencies("aabaacdb")
}
