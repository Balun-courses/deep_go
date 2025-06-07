package main

import (
	"fmt"
	"reflect"
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

type Container struct {
	constructors map[string]interface{}
}

func NewContainer() *Container {
	return &Container{
		constructors: make(map[string]interface{}),
	}
}

func (c *Container) RegisterType(name string, constructor interface{}) {
	if reflect.TypeOf(constructor).Kind() != reflect.Func {
		panic(fmt.Sprintf("constructor for %s must be a function", name))
	}
	c.constructors[name] = constructor
}

func (c *Container) Resolve(name string) (interface{}, error) {
	constructor, exists := c.constructors[name]
	if !exists {
		return nil, fmt.Errorf("dependency %s not found", name)
	}
	constructorType := reflect.TypeOf(constructor)
	if constructorType.NumOut() == 0 {
		return nil, fmt.Errorf("constructor for %s must return at least one value", name)
	}
	results := reflect.ValueOf(constructor).Call(nil)
	if len(results) == 0 {
		return nil, fmt.Errorf("constructor for %s returned no values", name)
	}
	for _, result := range results {
		if result.Type().Implements(reflect.TypeOf((*error)(nil)).Elem()) && !result.IsNil() {
			return nil, result.Interface().(error)
		}
	}

	return results[0].Interface(), nil
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
