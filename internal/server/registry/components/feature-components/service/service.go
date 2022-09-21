package service

import "test/internal/features/test-domain"

type Container struct {
	Domain1Service test_domain.Service
}

func NewContainer() *Container {
	return &Container{}
}
