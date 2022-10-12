package feature_components

import (
	"github.com/ihatiko/di"
	"test/internal/features/domain1"
	domain1api "test/internal/features/domain1/delivery/api"
	domain1service "test/internal/features/domain1/service"
	"test/internal/features/domain2"
	domain2api "test/internal/features/domain2/delivery/api"
	domain3api "test/internal/features/domain2/delivery/api"
	domain2service "test/internal/features/domain2/service"
	domain3service "test/internal/features/domain2/service"
	"test/internal/features/domain3"
)

func Registry() {
	SetRepository()
	SetService()
	SetDelivery()
}

func SetDelivery() {
	di.ProvideInterface[domain1.ApiHandler](domain1api.NewApiHandler)
	di.ProvideInterface[domain2.ApiHandler](domain2api.NewApiHandler)
	di.ProvideInterface[domain3.ApiHandler](domain3api.NewApiHandler)
}

func SetRepository() {

}

func SetService() {
	di.ProvideInterface[domain1.Service](domain1service.NewDomain1Service)
	di.ProvideInterface[domain2.Service](domain2service.NewDomain1Service)
	di.ProvideInterface[domain3.Service](domain3service.NewDomain1Service)
}
