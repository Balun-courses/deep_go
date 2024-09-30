package main

type Nothing interface{}

func main() {
	var nothing Nothing
	var empty interface{}
	var any any

	nothing = empty
	nothing = any

	empty = nothing
	empty = any

	any = nothing
	any = empty
}
