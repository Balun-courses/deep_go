package main

import (
	"fmt"
	"unsafe"
)

type eface struct {
	typ  unsafe.Pointer
	data unsafe.Pointer
}

type hmap struct {
	count     int
	flags     uint8
	B         uint8
	noverflow uint16
	hash0     uint32

	buckets    unsafe.Pointer
	oldbuckets unsafe.Pointer
	nevacuate  uintptr
}

func main() {
	data := make(map[int]int)
	iface := any(data)

	ifacePtr := unsafe.Pointer(&iface)
	emptyIfacePtr := (*eface)(ifacePtr)

	hm := (*hmap)(emptyIfacePtr.data)
	fmt.Println("count:", hm.count, "buckets:", 1<<hm.B)

	for i := 0; i < 500; i++ {
		data[i] = i * 2
	}

	fmt.Println("count:", hm.count, "buckets:", 1<<hm.B)

	for key := range data {
		delete(data, key)
	}

	fmt.Println("count:", hm.count, "buckets:", 1<<hm.B)
}
