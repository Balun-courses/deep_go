package main

import "strings"

func firstWay() {
	old := strings.Builder{}
	// manipulating with old..
	new := old
	_ = new
}

func secondWay() {
	old := strings.Builder{}
	// manipulating with old..
	new := strings.Builder{}
	new.WriteString(old.String())
	_ = new
}
