package main

type Integer1 interface {
	int | int16 | int32 | int64
}

type Integer2 interface {
	~int | ~int16 | ~int32 | ~int64
}

func Process1[T1 Integer1](value T1) {
	// ..
}

func Process11[T2 interface{ int | int16 | int32 | int64 }](value T2) {
	// ..
}

func Process12[T2 int | int16 | int32 | int64](value T2) { // interface{} can be omitted
	// ..
}

func Process2[T2 Integer2](value T2) {
	// ..
}

type MyInt1 int
type MyInt2 MyInt1

type MyIntAlias1 = int
type MyIntAlias2 = MyIntAlias1

func main() {
	Process1(int(100))
	Process1(MyInt1(100))
	Process1(MyInt2(100))
	Process1(MyIntAlias1(100))
	Process1(MyIntAlias2(100))

	Process2(int(100))
	Process2(MyInt1(100))
	Process2(MyInt2(100))
	Process1(MyIntAlias1(100))
	Process1(MyIntAlias2(100))
}
