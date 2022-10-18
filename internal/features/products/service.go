package products

import "context"

type Service interface {
	Domain1Get(ctx context.Context) error
	Domain1Post(ctx context.Context) error
	Domain1Patch(ctx context.Context) error
	Domain1Delete(ctx context.Context) error
}
