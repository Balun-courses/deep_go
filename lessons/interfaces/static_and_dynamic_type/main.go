package main

import (
	"time"
)

type Stringer interface {
	String() string
}

func main() {
	var s Stringer  // static type
	s = time.Time{} // dynamic type
	_ = s
}
