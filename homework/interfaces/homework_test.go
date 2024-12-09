package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type UserService struct {
	// not need to implement
	NotEmptyStruct bool
}
type MessageService struct {
	// not need to implement
	NotEmptyStruct bool
}

const (
	transient = iota
	singleton
)

type containerItem struct {
	ctor     interface{}
	lifetime int
	instance interface{}
}

type Container struct {
	m map[string]*containerItem
}

func NewContainer() *Container {
	return &Container{m: map[string]*containerItem{}}
}

func (c *Container) RegisterType(name string, constructor interface{}) {
	c.m[name] = &containerItem{ctor: constructor, lifetime: transient}
}

func (c *Container) Resolve(name string) (interface{}, error) {
	item, ok := c.m[name]

	if !ok {
		return nil, fmt.Errorf("type %s not registered", name)
	}

	fn, ok := item.ctor.(func() interface{})
	if !ok {
		return nil, fmt.Errorf("invalid constructor for type %s", name)
	}

	if item.lifetime == transient {
		return fn(), nil
	}

	// not thread-safety
	if item.instance == nil {
		item.instance = fn()
	}
	return item.instance, nil
}

func (c *Container) RegisterSingletonType(name string, constructor interface{}) {
	c.m[name] = &containerItem{ctor: constructor, lifetime: singleton}
}

func TestDIContainer(t *testing.T) {
	container := NewContainer()
	container.RegisterType("UserService", func() interface{} {
		return &UserService{}
	})
	container.RegisterType("MessageService", func() interface{} {
		return &MessageService{}
	})

	userService1, err := container.Resolve("UserService")
	assert.NoError(t, err)
	userService2, err := container.Resolve("UserService")
	assert.NoError(t, err)

	u1 := userService1.(*UserService)
	u2 := userService2.(*UserService)
	assert.False(t, u1 == u2)

	messageService, err := container.Resolve("MessageService")
	assert.NoError(t, err)
	assert.NotNil(t, messageService)

	paymentService, err := container.Resolve("PaymentService")
	assert.Error(t, err)
	assert.Nil(t, paymentService)

	container.RegisterSingletonType("UserService2", func() interface{} {
		return &UserService{}
	})
	userService3, err := container.Resolve("UserService2")
	assert.NoError(t, err)
	userService4, err := container.Resolve("UserService2")
	assert.NoError(t, err)
	u3 := userService3.(*UserService)
	u4 := userService4.(*UserService)
	assert.True(t, u3 == u4)
}
