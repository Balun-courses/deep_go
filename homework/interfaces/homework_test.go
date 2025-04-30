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
type Db struct {
	// not need to implement
	NotEmptyStruct bool
}

type Container struct {
	Services          map[string]any
	SingleTonServices map[string]bool
}

func NewContainer() *Container {
	return &Container{
		Services:          map[string]any{},
		SingleTonServices: map[string]bool{},
	}
}

func (c *Container) RegisterType(name string, constructor any) {
	c.Services[name] = constructor
}

func (c *Container) RegisterSingletonType(name string, constructor any) {
	c.Services[name] = constructor
	c.SingleTonServices[name] = true
}

func (c *Container) Resolve(name string) (any, error) {
	value, found := c.Services[name]
	if !found {
		return nil, fmt.Errorf("service %s not found", name)
	}

	service, ok := value.(func() any)
	if !ok {
		if c.SingleTonServices[name] {
			return c.Services[name], nil
		}

		return nil, fmt.Errorf("service %s not a constructor", name)
	}

	if c.SingleTonServices[name] {
		c.Services[name] = service()
		return c.Services[name], nil
	}

	return service(), nil
}

func TestDIContainer(t *testing.T) {
	container := NewContainer()

	container.RegisterSingletonType("Db", func() interface{} {
		return &Db{}
	})

	db1, err := container.Resolve("Db")
	assert.NoError(t, err)

	db2, err := container.Resolve("Db")
	assert.NoError(t, err)

	assert.True(t, db1 == db2)

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
