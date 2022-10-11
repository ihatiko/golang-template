package feature_components

import (
	"github.com/ihatiko/di"
	test_domain "test/internal/features/test-domain"
	"test/internal/features/test-domain/delivery/api"
	"test/internal/features/test-domain/service"
)

func Registry() {
	SetRepository()
	SetService()
	SetDelivery()
}

func SetDelivery() {
	di.ProvideInterface[test_domain.ApiHandler](api.NewApiHandler)
}

func SetRepository() {

}

func SetService() {
	di.ProvideInterface[test_domain.Service](service.NewDomain1Service)
}
