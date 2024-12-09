package main

type Creator[T any] func() T

func NewInstance[T any](creator Creator[T]) T {
	return creator()
}

type (
	Book struct{}
	Game struct{}
)

func main() {
	bookCreator := func() Book { return Book{} }
	gameCreator := func() Game { return Game{} }

	book := NewInstance(bookCreator)
	_ = book

	game := NewInstance(gameCreator)
	_ = game
}
