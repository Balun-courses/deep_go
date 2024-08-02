package main

import "fmt"

func mult(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}

func main() {
	fmt.Println(multOrDivide(true, mult, divide, 17, 3)) //здесь вызываем наши "эмулирующую" ленивые вычисления функцию, которая через 1 аргумент решает,
	// пора ли уже вызывать вычисления функции или нет
	fmt.Println(multOrDivide(false, mult, divide, 17, 3))
}

// наш if - else откладывает выполнение наших "одиночных" функций
func multOrDivide(add bool, onMult, onDivide func(t, z int) int, t, z int) int {
	if add {
		return onMult(t, z)
	}
	return onDivide(t, z)
}
