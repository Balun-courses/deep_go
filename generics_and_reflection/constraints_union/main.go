package main

type _ interface {
	[]int | comparable // compilation error
}

type _ interface {
	string | error // compilation error
}

type _ interface {
	string | interface{ Do() } // compilation error
}

type _ interface {
	string | interface{ int } // ok
}

type _ interface {
	string | interface{} // ok
}
