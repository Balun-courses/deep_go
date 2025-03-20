package main

type Getter interface {
	Get() interface{}
}

type PointerGetter struct{}

func (p *PointerGetter) Get() interface{} {
	return nil
}

func get() PointerGetter {
	return PointerGetter{}
}

func main() {
	var ptrGetter PointerGetter
	ptrGetter.Get() // ok

	var getter Getter = get() // error
	_ = getter
}
