package service

import (
	"context"
	"github.com/opentracing/opentracing-go"
)

type TestDomain struct {
}

func NewDomain1Service() *TestDomain {
	return &TestDomain{}
}

func (s TestDomain) Domain1Get(ctx context.Context) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "service.Domain1Get")
	span.SetTag("user-agent", "some-user-agent")
	defer span.Finish()
	return nil
}

func (s TestDomain) Domain1Post(ctx context.Context) error {
	return nil
}

func (s TestDomain) Domain1Patch(ctx context.Context) error {
	return nil
}

func (s TestDomain) Domain1Delete(ctx context.Context) error {
	return nil
}
