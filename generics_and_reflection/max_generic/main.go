package main

// constraint
type Signed interface {
	int | int8 | int16 | int32 | int64
}

// constraint
type Unsigned interface {
	uint | uint8 | uint16 | uint32 | uint64
}

// constraint
type Integer interface {
	Signed | Unsigned
}

func Max[T Integer](lhs, rhs T) T {
	if lhs >= rhs {
		return lhs
	} else {
		return rhs
	}
}

func main() {
	var lhs1, rhs1 int = 10, 20
	_ = Max[int](lhs1, rhs1)

	var lhs2, rhs2 uint = 10, 20
	_ = Max[uint](lhs2, rhs2)
}
