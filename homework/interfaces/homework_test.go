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

type Container struct {
	services map[string]Service
}

type Service interface{}

func NewContainer() *Container {
	// need to implement
	return &Container{
		services: make(map[string]Service),
	}
}

func (c *Container) RegisterType(name string, constructor interface{}) {
	fmt.Printf("CONTEINER %+v %+v\n", c, constructor)
	c.services[name] = Service(constructor)
	fmt.Printf("CONTEINER %+v\n", c)

}

func (c *Container) Resolve(name string) (interface{}, error) {
	var err error
	service := c.services[name]

	if service != nil {
		fn, ok := service.(func() interface{})
		if !ok {
			err = fmt.Errorf("service is not a function")
			return nil, err
		}
		return fn(), nil
	}
	err = fmt.Errorf("service not found")

	return nil, err
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
