package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ihatiko/log"
	"github.com/opentracing/opentracing-go"
	"test/internal/features/products"
	"test/internal/features/products/models"
)

type Handler struct {
	Service products.Service
}

func NewApiHandler(service products.Service) *Handler {
	return &Handler{Service: service}
}

func (h Handler) TestGet(context *fiber.Ctx) error {
	span, ctx := opentracing.StartSpanFromContext(context.UserContext(), "products.TestGet")
	defer span.Finish()
	h.Service.Domain1Get(ctx)
	context.Send([]byte("ok"))
	return nil
}

func (h Handler) TestPost(context *fiber.Ctx) error {
	span, _ := opentracing.StartSpanFromContext(context.UserContext(), "products.TestPost")
	defer span.Finish()
	log.Info("POST")
	str := &models.TestDomainResponse{
		Field1: "1",
		Field2: "2",
		Field3: "3",
	}

	return context.JSON(str)
}

func (h Handler) TestPatch(context *fiber.Ctx) error {
	context.Send([]byte("ok"))
	return nil
}

func (h Handler) TestDelete(context *fiber.Ctx) error {
	context.Send([]byte("ok"))
	return nil
}
