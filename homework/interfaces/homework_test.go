package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
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

type constructorFunc func() interface{}

type Container struct {
	mu           sync.RWMutex
	constructors map[string]constructorFunc
}

func NewContainer() *Container {
	return &Container{
		constructors: make(map[string]constructorFunc),
	}
}

func (c *Container) RegisterType(name string, constructor interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if fn, ok := constructor.(func() interface{}); ok {
		c.constructors[name] = fn
	}
}

func (c *Container) Resolve(name string) (interface{}, error) {
	c.mu.RLock()
	constructor, ok := c.constructors[name]
	c.mu.RUnlock()

	if !ok {
		return nil, errors.New("type not registered: " + name)
	}
	return constructor(), nil
}

func (c *Container) RegisterSingletonType(name string, constructor interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if fn, ok := constructor.(func() interface{}); ok {
		var (
			once     sync.Once
			instance interface{}
		)
		c.constructors[name] = func() interface{} {
			once.Do(func() {
				instance = fn()
			})
			return instance
		}
	}
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
