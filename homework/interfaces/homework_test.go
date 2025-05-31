package main

import (
	"fmt"
	"sync"
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

type constructorFunc func() interface{}

type Container struct {
	mu           sync.RWMutex
	constructors map[string]constructorFunc
	singletons   map[string]interface{}
	isSingleton  map[string]bool
}

func NewContainer() *Container {
	return &Container{
		constructors: make(map[string]constructorFunc),
		singletons:   make(map[string]interface{}),
		isSingleton:  make(map[string]bool),
	}
}

func (c *Container) RegisterType(name string, constructor interface{}) {
	c.register(name, constructor, false)
}

func (c *Container) RegisterSingletonType(name string, constructor interface{}) {
	c.register(name, constructor, true)
}

func (c *Container) register(name string, constructor interface{}, singleton bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	fn, ok := constructor.(func() interface{})
	if !ok {
		panic(fmt.Sprintf("constructor for %s must be of type func() interface{}", name))
	}
	c.constructors[name] = fn
	c.isSingleton[name] = singleton
}

func (c *Container) Resolve(name string) (interface{}, error) {
	c.mu.RLock()
	constructor, exists := c.constructors[name]
	isSingleton := c.isSingleton[name]
	c.mu.RUnlock()

	if !exists {
		return nil, fmt.Errorf("type %s not registered in container", name)
	}

	if isSingleton {
		c.mu.RLock()
		instance, singletonExists := c.singletons[name]
		c.mu.RUnlock()
		if singletonExists {
			return instance, nil
		}

		c.mu.Lock()
		defer c.mu.Unlock()

		if instance, singletonExists = c.singletons[name]; singletonExists {
			return instance, nil
		}
		instance = constructor()
		c.singletons[name] = instance
		return instance, nil
	}

	return constructor(), nil
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
