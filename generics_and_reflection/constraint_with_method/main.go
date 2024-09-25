package main

type Stringer interface {
	String() string
}

func ToString[T Stringer](values ...T) []string {
	result := make([]string, 0, len(values))
	for idx := range values {
		result = append(result, values[idx].String())
	}

	return result
}

type Data struct{}

func (d Data) String() string {
	return "string"
}

func main() {
	data := Data{}
	var idata Stringer = data

	_ = ToString(data)
	_ = ToString(idata)
	_ = ToString(100)
}
