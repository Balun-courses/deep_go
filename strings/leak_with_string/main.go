package main

func FindData(filename string) string {
	var data string
	// reading data from file..

	for i := 0; i < len(data)-1; i++ {
		if data[i] == '\n' && data[i+1] == '\t' {
			return data[i+2 : i+22] // capacity leak
		}
	}

	return ""
}

func main() {
	_ = FindData("data.txt")
	// potentially high memory consumption
}
