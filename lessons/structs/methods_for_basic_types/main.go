package main

type Age int

func (age Age) LargerThan(other Age) bool {
	return age > other
}

type FilterFunc func(int) bool

func (ff FilterFunc) Filter(value int) bool {
	return ff(value)
}

type StringSet map[string]struct{}

func (ss StringSet) Has(key string) bool {
	_, found := ss[key]
	return found
}

func (ss StringSet) Add(key string) {
	ss[key] = struct{}{}
}

func (ss StringSet) Remove(key string) {
	delete(ss, key)
}
