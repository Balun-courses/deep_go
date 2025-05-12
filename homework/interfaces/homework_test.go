package main

import (
	"errors"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type UserService struct {
	NotEmptyStruct bool
}

type MessageService struct {
	NotEmptyStruct bool
}

type Container struct {
	services sync.Map // хранилище зарегистрированных конструкторов
}

func NewContainer() *Container {
	return &Container{}
}

func (c *Container) RegisterType(name string, constructor interface{}) {
	c.services.Store(name, constructor)
}

func (c *Container) Resolve(name string) (interface{}, error) {
	// Получаем конструктор из хранилища
	val, ok := c.services.Load(name)
	if !ok {
		return nil, errors.New("service not found: " + name)
	}

	// Проверяем тип конструктора
	constructor, ok := val.(func() interface{})
	if !ok {
		return nil, errors.New("invalid constructor for service: " + name)
	}

	// Создаем новый экземпляр
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
