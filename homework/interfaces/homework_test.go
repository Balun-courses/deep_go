package main

import (
	"testing"

	"fmt"
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

type Container struct {
	dependency map[string]interface{}
}

func NewContainer() *Container {
	return &Container{
		dependency: make(map[string]interface{}),
	}
}

func (c *Container) RegisterType(name string, constructor interface{}) {
	if c.dependency == nil {
		return
	}
	c.dependency[name] = constructor
}

func (c *Container) Resolve(name string) (interface{}, error) {
	constructor, ok := c.dependency[name]
	if !ok {
		return nil, fmt.Errorf("no constructor registered for %s", name)
	}
	fun, ok := constructor.(func() interface{})
	if !ok {
		return nil, fmt.Errorf("constructor does not implement %T", constructor)
	}

	return fun(), nil
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
}
