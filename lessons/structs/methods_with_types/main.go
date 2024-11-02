package main

import "io"

type (
	IntAlias = int
	IntType  int
)

func (i int) Do()      {} // compilation error: basic type
func (i IntAlias) Do() {} // compilation error: basic type
func (i IntType) Do()  {} // ok

type (
	IntPtrAlias = *int
	IntPtrType  *int
)

func (i IntPtrAlias) Do() {} // compilation error: pointer
func (i IntPtrType) Do()  {} // compilation error: pointer

type (
	CloserAlias = io.Closer
	CloserType  io.Closer
)

func (i io.Closer) Do()   {} // compilation error: interface
func (i CloserAlias) Do() {} // compilation error: interface
func (i CloserType) Do()  {} // compilation error: interface
