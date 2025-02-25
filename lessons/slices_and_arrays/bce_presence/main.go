package main

// go run -gcflags="-d=ssa/check_bce" main.go

func fn1(s []int) {
	for i := range s {
		_ = s[i]
		_ = s[i:]
		_ = s[:i+1]
	}
}

func fn2(s []int) {
	for i := 0; i < len(s); i++ {
		_ = s[i]
		_ = s[i:]
		_ = s[:i+1]
	}
}

func main() {}
