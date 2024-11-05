package main

type BSTItem interface {
	Less(BSTItem) bool
}

type BinarySearhTree struct {
	item  BSTItem
	left  *BinarySearhTree
	right *BinarySearhTree
}
