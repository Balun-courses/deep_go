package main

func process1[T1 int, T2 *T1]() {
}

func process2[T1 *T2, T2 int]() {
}

func process3[T1 *T1]() {
}

func process4[T ~string | ~int, A ~[2]T, B ~chan T]() {

}
