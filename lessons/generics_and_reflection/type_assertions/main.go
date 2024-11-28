package main

import "fmt"

func IsInt1[T any](x T) bool {
	_, ok := any(x).(int)
	return ok
}

func Is1[T int | string](x T) {
	switch any(x).(type) {
	//case T:
	//	fmt.Println("int")
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	}
}

func IsInt2[T any](x T) bool {
	_, ok := x.(int) // cannot use type assertion on type parameter
	return ok
}

func Is2[T int | string](x T) {
	switch x.(type) { // cannot use type switch on type parameter
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	}
}

func main() {
	fmt.Println("IsInt(100):", IsInt1(100))
	fmt.Println(`IsInt("100"):`, IsInt1("100"))

	Is1("100")
	Is1(100)
}
