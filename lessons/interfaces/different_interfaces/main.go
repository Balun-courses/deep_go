package main

type Stringer interface {
	String() string
}

// Embeds all types whose underlying type is []byte
type AnyByteSlice interface {
	~[]byte
}

// Embeds type union
type Unsigned interface {
	uint | uintptr | uint8 | uint16 | uint32 | uint64
}

func main() {
	var stringer Stringer
	_ = stringer

	var any AnyByteSlice
	_ = any

	var unsigned Unsigned
	_ = unsigned
}
