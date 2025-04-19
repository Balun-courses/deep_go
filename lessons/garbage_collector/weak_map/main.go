package main

import (
	"runtime"
	"time"
	"weak"
)

type WeakMap struct {
	data map[string]weak.Pointer[string]
}

func NewWeakMap() WeakMap {
	return WeakMap{
		data: make(map[string]weak.Pointer[string]),
	}
}

func (w *WeakMap) Set(key string, value *string) {
	runtime.AddCleanup(value, w.Delete, key)
	w.data[key] = weak.Make(value)
}

func (w *WeakMap) Get(key string) *string {
	if ptr, ok := w.data[key]; ok {
		return ptr.Value()
	}

	return nil
}

func (w *WeakMap) Delete(key string) {
	delete(w.data, key)
}

func main() {
	data := NewWeakMap()

	key := "my key"
	value := "my data"
	data.Set(key, &value)

	runtime.GC()
	time.Sleep(time.Second)

	pointer := data.Get(key)
	println(pointer == nil)
}
