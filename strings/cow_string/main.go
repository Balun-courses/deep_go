package main

// Copy-On-Write
type COWString struct {
	data []byte
	refs *int
}

func NewString(values ...byte) COWString {
	return COWString{
		data: values,
		refs: new(int),
	}
}

func (s *COWString) Length() int {
	return len(s.data)
}

func (s *COWString) Capacity() int {
	return cap(s.data)
}

func (s *COWString) ToString() string {
	return string(s.data) // copying...
}

func (s *COWString) Get(idx int) byte {
	// without bound checking
	return s.data[idx]
}

func (s *COWString) Set(idx int, value byte) {
	// without bound checking

	if *s.refs > 0 {
		s.data = append([]byte(nil), s.data...)
	}

	s.data[idx] = value
}

func (s *COWString) Copy() COWString {
	(*s.refs)++
	return COWString{
		data: s.data,
		refs: s.refs,
	}
}

func (s *COWString) Append(values ...byte) {
	if *s.refs > 0 {
		s.data = append([]byte(nil), s.data...)
	}

	s.data = append(s.data, values...)
}

func main() {
	str := NewString([]byte("Hello world")...)
	temp := str
	_ = temp
}
