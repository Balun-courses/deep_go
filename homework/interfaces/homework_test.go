package main

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"sync"
	"testing"
)

type UserService struct {
	NotEmptyStruct bool
}
type MessageService struct {
	NotEmptyStruct bool
}

type Container struct {
	mu           sync.RWMutex
	constructors map[string]interface{}
}

func NewContainer() *Container {
	return &Container{
		constructors: make(map[string]interface{}),
	}
}

func (c *Container) RegisterType(name string, constructor interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if constructor == nil {
		panic("constructor cannot be nil")
	}
	c.constructors[name] = constructor
}

func (c *Container) Resolve(name string) (interface{}, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	constructor, ok := c.constructors[name]
	if !ok {
		return nil, fmt.Errorf("type '%s' not registered", name)
	}

	ctorVal := reflect.ValueOf(constructor)

	if ctorVal.Kind() != reflect.Func || ctorVal.Type().NumIn() != 0 || ctorVal.Type().NumOut() != 1 {
		return nil, errors.New("constructor must be a function with no input arguments and one return value")
	}

	result := ctorVal.Call(nil)
	return result[0].Interface(), nil
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
