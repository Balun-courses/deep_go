package main

import (
	"fmt"
	"unsafe"
)

type emptyInterface struct {
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
	emptyIfacePtr := (*emptyInterface)(ifacePtr)

	hm := (*hmap)(emptyIfacePtr.data)
	fmt.Println(hm.count, 1<<hm.B)

	for i := 0; i < 500; i++ {
		data[i] = i * 2
	}

	fmt.Println(hm.count, 1<<hm.B)

	for key := range data {
		delete(data, key)
	}

	fmt.Println(hm.count, 1<<hm.B)
}
