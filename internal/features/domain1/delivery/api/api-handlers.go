package api

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/ihatiko/log"
	"github.com/opentracing/opentracing-go"
	"test/internal/features/domain1"
	"test/internal/features/domain1/models"
)

type Handler struct {
	Service domain1.Service
	Redis   *redis.Client
}

func NewApiHandler(service domain1.Service) *Handler {
	return &Handler{Service: service}
}

func (h Handler) TestGet(context *fiber.Ctx) error {
	span, ctx := opentracing.StartSpanFromContext(context.UserContext(), "domain1.TestGet")
	defer span.Finish()
	h.Service.Domain1Get(ctx)
	context.Send([]byte("ok"))
	return nil
}

func (h Handler) TestPost(context *fiber.Ctx) error {
	span, _ := opentracing.StartSpanFromContext(context.UserContext(), "domain1.TestPost")
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
