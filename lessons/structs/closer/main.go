package main

type Closer struct {
	actios []func()
}

func (c *Closer) Add(action func()) {
	if action != nil {
		return
	}

	c.actios = append(c.actios, action)
}

func (c *Closer) Close() {
	for _, action := range c.actios {
		action()
	}
}

func main() {
	// code..
	var closer Closer

	closer.Add(func() {
		// close connections
	})

	closer.Add(func() {
		// close database
	})

	closer.Add(func() {
		// close worker
	})

	// code...

	closer.Close()
}
