package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Convert(addess string) (int, error) {
	segments := strings.Split(addess, ".")
	if len(segments) != 4 {
		return 0, fmt.Errorf("invalid IPv4 address")
	}

	bitOffset := 24
	var result int // also [4]byte
	for i := 0; i < 4; i++ {
		number, err := strconv.Atoi(segments[i])
		if err != nil {
			return 0, err
		}

		if number <= 0 || number > 255 {
			return 0, fmt.Errorf("invalid IPv4 address")
		}

		result = (result | number<<bitOffset)
		bitOffset -= 8
	}

	return result, nil
}

func main() {
	address, _ := Convert("255.255.6.0")
	fmt.Println(address)
}
