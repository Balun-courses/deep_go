package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

type FileWrapper struct {
	file *os.File
}

func NewFileWrapper(path string) *FileWrapper {
	file, _ := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0666)

	ptr := &FileWrapper{file: file}

	runtime.AddCleanup(ptr, func(f *os.File) { // replace with FileWrapper
		fmt.Println("Closing file:", f.Name())
		_ = f.Close()
	}, ptr.file)

	return ptr
}

func main() {
	obj := NewFileWrapper(os.DevNull)
	_ = obj
	obj = nil

	runtime.GC()
	time.Sleep(time.Second)

	fmt.Println("Object is collected by GC")
}
