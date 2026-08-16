package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type RealConstructor = func() any

type UserService struct {
	// not need to implement
	NotEmptyStruct bool
}
type MessageService struct {
	// not need to implement
	NotEmptyStruct bool
}

type Container struct {
	// need to implement
	ifs map[string]any
}

func NewContainer() *Container {
	// need to implement
	return &Container{make(map[string]any)}
}

func (c *Container) RegisterType(name string, constructor interface{}) {
	if c.ifs == nil {
		c.ifs = make(map[string]any)
	}

	c.ifs[name] = constructor
}

func (c *Container) Resolve(name string) (interface{}, error) {
	if constructor, ok := c.ifs[name]; ok {
		if realConstructor, ok := constructor.(RealConstructor); ok {
			return realConstructor(), nil
		}
		return nil, fmt.Errorf("%s constructor has invalid prototype", name)
	}

	return nil, fmt.Errorf("name not resolved: %s", name)
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
