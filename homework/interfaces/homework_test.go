package main

import (
	"errors"
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
	services map[string]func() interface{}
}

// создать DI контейнер
func NewContainer() *Container {
	// need to implement
	return &Container{
		services: make(map[string]func() interface{}),
	}
}

// зарегистрировать конструктор по созданию типа
func (c *Container) RegisterType(name string, constructor interface{}) {
	init, valid := constructor.(func() interface{})
	if !valid {
		return
	}
	c.services[name] = init
}

// создать объект с использованием конструктора
func (c *Container) Resolve(name string) (interface{}, error) {
	init, valid := c.services[name]
	if !valid {
		return nil, errors.New("no constructor is registered for " + name)
	}
	return init(), nil
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
