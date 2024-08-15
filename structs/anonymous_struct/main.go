package main

func main() {
	var book = struct {
		author struct {
			name    string
			surname string
		}
		title string
	}{
		author: struct {
			name    string
			surname string
		}{
			name:    "Name",
			surname: "Surname",
		},
		title: "Title",
	}

	_ = book
}
