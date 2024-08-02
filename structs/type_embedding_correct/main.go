package main

import "io"

type LoggerV1 struct {
	closer io.WriteCloser
}

func (l *LoggerV1) Write(data []byte) (int, error) {
	return l.closer.Write(data)
}

func (l *LoggerV1) Close() error {
	return l.closer.Close()
}

type LoggerV2 struct {
	io.WriteCloser
}

func main() {
	v1 := LoggerV1{}
	v1.Write(nil)
	v1.Close()

	v2 := LoggerV2{}
	v2.Write(nil)
	v2.Close()
}
