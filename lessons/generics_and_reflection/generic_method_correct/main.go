package main

// constraint
type Signed interface {
	int | int8 | int16 | int32 | int64
}

type Counter struct {
	value int64
}

func Add[T Signed](c *Counter, value T) {
	c.value += int64(value)
}
