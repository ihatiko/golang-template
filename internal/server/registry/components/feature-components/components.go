package feature_components

import (
	"github.com/ihatiko/di"
	"test/internal/features/basket"
	basketApi "test/internal/features/basket/delivery/api"
	basketService "test/internal/features/basket/service"
	"test/internal/features/payments"
	paymentsApi "test/internal/features/payments/delivery/api"
	paymentsService "test/internal/features/payments/service"
	"test/internal/features/products"
	productsApi "test/internal/features/products/delivery/api"
	productsRepository "test/internal/features/products/repository"
	productsService "test/internal/features/products/service"
)

func Registry() {
	SetRepository()
	SetService()
	SetDelivery()
}

func SetDelivery() {
	di.ProvideInterface[products.ApiHandler](productsApi.NewApiHandler)
	di.ProvideInterface[payments.ApiHandler](paymentsApi.NewApiHandler)
	di.ProvideInterface[basket.ApiHandler](basketApi.NewApiHandler)
}

func SetRepository() {
	di.ProvideInterface[products.Repository](productsRepository.NewProductsRepository)
}

func SetService() {
	di.ProvideInterface[products.Service](productsService.NewProductsService)
	di.ProvideInterface[payments.Service](paymentsService.NewPaymentsService)
	di.ProvideInterface[basket.Service](basketService.NewBasketService)
}
