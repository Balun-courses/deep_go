package main

type client struct {
	name    string
	surname string
}

func updateSurname(data map[int]client, id int, surname string) {
	object := data[id] // copy
	object.surname = surname
	data[id] = object // copy

	// data[id].surname = surname -> compilation error
}

func updateSurnameByPointer(data map[int]*client, id int, surname string) {
	data[id].surname = surname
}

func updateSurnamesSlice(data map[int][]string, id int, surname string) {
	data[id][5] = surname
}
