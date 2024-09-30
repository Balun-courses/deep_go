package main

import "fmt"

func main() {
	value := func() (result int) {
		defer func() {
			v := recover()
			result = v.(int)
		}()

		func() {
			func() {
				func() {
					panic(123)
					// ...
				}()
				// ...
			}()
			// ...
		}()

		return
	}()

	fmt.Println(value)
}
