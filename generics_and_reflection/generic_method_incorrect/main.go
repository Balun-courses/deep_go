package main

// constraint
type Signed interface {
	int | int8 | int16 | int32 | int64
}

type Counter struct {
	value int64
}

func (c *Counter) Add[T Signed](value T) { // compilation erro
	c.value += int64(value)
}