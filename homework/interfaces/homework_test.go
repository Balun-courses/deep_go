package main

import (
	"errors"
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

type Container struct {
	typeConstructor map[string]interface{}
}

func NewContainer() *Container {
	return &Container{typeConstructor: make(map[string]interface{})}
}

func (c *Container) RegisterType(name string, constructor interface{}) {
	c.typeConstructor[name] = constructor
}

func (c *Container) Resolve(name string) (obj interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			obj = nil
			err = fmt.Errorf("failed to run constructor for %s: %v", name, r)
		}
	}()
	constructor, ok := c.typeConstructor[name]
	if !ok {
		return nil, errors.New("no constructor registered")
	}
	return constructor.(func() interface{})(), nil
}

func TestDIContainer(t *testing.T) {
	container := NewContainer()
	container.RegisterType("UserService", func() interface{} {
		return &UserService{}
	})
	container.RegisterType("MessageService", func() interface{} {
		return &MessageService{}
	})
	container.RegisterType("int", 0)

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

	intService, err := container.Resolve("int")
	assert.Error(t, err)
	assert.Nil(t, intService)
}
