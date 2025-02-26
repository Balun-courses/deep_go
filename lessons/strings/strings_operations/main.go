package main

func writeToStringSymbol() {
	str := "hello"
	// str[0] = 'H' -> compilation error
	_ = str
}

func rewriteString() {
	str := "hello"
	str = "world"
	// str[0] = 'H' -> compilation error
	_ = str
}

func takeStringSymbolAddress() {
	str := "hello"
	// pointer := &str[0] -> compilation error
	_ = str
}

func slicingString() {
	/*
		const str = "hello"
		_ = str[0:10] -> compilation error
	*/
}
