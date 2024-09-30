package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type stateFn func(*lexer) stateFn

type lexer struct {
	*bufio.Reader
}

func run(l *lexer) {
	for state := lexText; state != nil; {
		state = state(l)
	}
}

var output = make(chan int)

func lexText(l *lexer) stateFn {
	for r, _, err := l.ReadRune(); err != io.EOF; r, _, err = l.ReadRune() {
		if '0' <= r && r <= '9' { //если попалась цифра
			l.UnreadRune()
			return lexNumber //переход состояния
		}
	}
	close(output)
	return nil // Стоп машина.
}

func lexNumber(l *lexer) stateFn {
	var s string
	for r, _, err := l.ReadRune(); err != io.EOF; r, _, err = l.ReadRune() {
		if '0' > r || r > '9' {
			num, _ := strconv.Atoi(s)
			output <- num  //передаем для утилизации
			return lexText //переход состояния
		}
		s += string(r)
	}
	num, _ := strconv.Atoi(s)
	output <- num
	close(output)
	return nil // Стоп машина.
}
func main() {
	var sum int
	a := "hell 3456 fgh 25 fghj 2128506 fgh 77"
	fmt.Println("Числа из строки: ", a)
	rr := strings.NewReader(a)
	lexy := lexer{
		bufio.NewReader(rr),
	}

	go run(&lexy)
	for nums := range output {
		fmt.Println(nums)
		sum += nums
	}
	fmt.Println("В сумме дают: ", sum)
}
