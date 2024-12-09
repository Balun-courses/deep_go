package main

type BSTItem interface {
	Less(BSTItem) bool
}

type BSTNode struct {
	item  BSTItem
	left  *BSTNode
	right *BSTNode
}
