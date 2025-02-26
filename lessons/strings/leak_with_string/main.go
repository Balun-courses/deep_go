package main

func findSequence(data string) string {
	for i := 0; i < len(data)-1; i++ {
		if data[i] == '\n' && data[i+1] == '\t' {
			return data[i+2 : i+22]
		}
	}

	return ""
}

func processBigData() {
	var data string
	// let's imagine that data was read from a file

	sequence := findSequence(data)
	_ = sequence // using of sequence later
}
