package main

// endianTypes — ограничение на поддерживаемые беззнаковые целочисленные типы.
type endianTypes interface {
	~uint16 | ~uint32 | ~uint64
}

// ToLittleEndian конвертирует число из Big Endian (прямого порядка
// следования байт) в Little Endian (обратный порядок следования байт).
// Реализовано только битовыми операциями и не зависит от конкретного типа:
// разрядность T определяется из его максимального значения (T(0) - 1).
func ToLittleEndian[T endianTypes](n T) T {
	width := 0
	for m := T(0) - 1; m != 0; m >>= 1 {
		width++
	}

	numBytes := width / 8
	var result T
	for i := 0; i < numBytes; i++ {
		b := (n >> (T(i) * 8)) & 0xFF
		result |= b << (T(numBytes-1-i) * 8)
	}
	return result
}
