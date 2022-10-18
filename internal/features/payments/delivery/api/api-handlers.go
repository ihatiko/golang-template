package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ihatiko/log"
	"test/internal/features/payments"
	"test/internal/features/products/models"
)

type Handler struct {
	Service payments.Service
}

func NewApiHandler(service payments.Service) *Handler {
	return &Handler{Service: service}
}

func (h Handler) TestGet(context *fiber.Ctx) error {
	context.Send([]byte("ok"))
	return nil
}

func (h Handler) TestPost(context *fiber.Ctx) error {
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
