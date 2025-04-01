package main

import "context"

type UserService interface {
	Save(ctx context.Context, name string) (int64, error)
	Get(ctx context.Context, id int64) (interface{}, error)
}

type UserRepository interface {
	Save(ctx context.Context, name string) (int64, error)
	Get(ctx context.Context, id int64) (interface{}, error)
}

type Provider struct{}

func NewProvider() Provider {
	return Provider{}
}

func (p *Provider) ProvideUserRepo() UserService
