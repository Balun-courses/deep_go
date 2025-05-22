package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

type UserService struct {
	NotEmptyStruct bool
}
type MessageService struct {
	NotEmptyStruct bool
}
type Singleton struct {
	NotEmptyStruct bool
}

type Container struct {
	dependencies map[string]interface{}
}

func NewContainer() *Container {
	return &Container{dependencies: make(map[string]interface{})}
}

func (c *Container) RegisterType(name string, constructor interface{}) {
	c.dependencies[name] = constructor
}

func (c *Container) RegisterSingletonType(name string, constructor interface{}) {
	if _, ok := c.dependencies[name]; ok {
		return
	}
	fn, ok := constructor.(func() interface{})
	if !ok {
		panic("constructor must be func() interface{} for singleton")
	}
	c.dependencies[name] = fn()
}

func (c *Container) Resolve(name string) (interface{}, error) {
	constructor, ok := c.dependencies[name]
	if !ok {
		return nil, errors.New("no constructor registered")
	}
	if fn, ok := constructor.(func() interface{}); ok {
		return fn(), nil
	}
	return constructor, nil
}

func TestDIContainer(t *testing.T) {
	container := NewContainer()
	container.RegisterType("UserService", func() interface{} {
		return &UserService{}
	})
	container.RegisterType("MessageService", func() interface{} {
		return &MessageService{}
	})
	container.RegisterSingletonType("Singleton", func() interface{} {
		return &Singleton{}
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

	service3, err := container.Resolve("Singleton")
	assert.NoError(t, err)
	service4, err := container.Resolve("Singleton")
	assert.NoError(t, err)
	s3 := service3.(*Singleton)
	s4 := service4.(*Singleton)
	assert.True(t, s3 == s4)
}
