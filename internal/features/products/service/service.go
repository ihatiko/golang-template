package service

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"test/internal/features/products"
)

type ProductsService struct {
	Repository products.Repository
}

func NewProductsService(repository products.Repository) *ProductsService {
	return &ProductsService{Repository: repository}
}

func (s ProductsService) Domain1Get(ctx context.Context) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.Domain1Get")
	span.SetTag("user-agent", "some-user-agent")
	defer span.Finish()
	return nil
}

func (s ProductsService) Domain1Post(ctx context.Context) error {
	return nil
}

func (s ProductsService) Domain1Patch(ctx context.Context) error {
	return nil
}

func (s ProductsService) Domain1Delete(ctx context.Context) error {
	return nil
}
