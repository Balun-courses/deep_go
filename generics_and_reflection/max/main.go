package main

func MaxInt(lhs, rhs int) int {
	if lhs >= rhs {
		return lhs
	} else {
		return rhs
	}
}

func MaxUInt(lhs, rhs uint) uint {
	if lhs >= rhs {
		return lhs
	} else {
		return rhs
	}
}

func main() {
	var lhs1, rhs1 int = 10, 20
	_ = MaxInt(lhs1, rhs1)

	var lhs2, rhs2 uint = 10, 20
	_ = MaxUInt(lhs2, rhs2)
}
