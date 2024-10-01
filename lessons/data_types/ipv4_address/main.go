package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Convert(address string) (uint32, error) {
	const octetsCount = 4
	segments := strings.Split(address, ".")
	if len(segments) != octetsCount {
		return 0, fmt.Errorf("invalid IPv4 address")
	}

	var result uint32 // also possible [4]byte
	for idx := 0; idx < octetsCount; idx++ {
		number, err := strconv.Atoi(segments[idx])
		if err != nil {
			return 0, err
		}

		if number < 0 || number > 255 {
			return 0, fmt.Errorf("invalid IPv4 address")
		}

		bitOffset := (octetsCount - idx - 1) * 8
		result |= uint32(number << bitOffset)
		fmt.Printf("%08b = %d\n", number, number)
	}

	return result, nil
}

func main() {
	address, _ := Convert("255.255.6.0")
	fmt.Printf("Address: %b = %d\n", address, address)
}
