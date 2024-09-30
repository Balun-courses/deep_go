package main

func IsSetBit(number, index int) bool {
	return (number & (1 << index)) != 0
}

func SetBit(number, index int) int {
	return number | (1 << index)
}

// TODO:
func InverseBit(number, index int) int {
	return number ^ (1 << index)
}

// TODO:
func ResetBit(number, index int) int {
	return number & ^(1 << index)
}
