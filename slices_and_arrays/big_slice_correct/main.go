package main

import "os"

func FindData(filename string) []byte {
	data, _ := os.ReadFile(filename)

	for i := 0; i < len(data)-1; i++ {
		if data[i] == 0xFF && data[i+1] == 0x00 {
			partData := make([]byte, 20)
			for j := i; j < 20; j++ {
				partData[j] = data[j]
			}

			return partData
		}
	}

	return nil
}

func main() {
	_ = FindData("data.bin")
}
