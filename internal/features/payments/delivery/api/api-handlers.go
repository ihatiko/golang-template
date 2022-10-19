package api

import (
	"github.com/gofiber/fiber/v2"
	"test/internal/features/payments"
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
	return nil
}

func (h Handler) TestPatch(context *fiber.Ctx) error {
	context.Send([]byte("ok"))
	return nil
}

func (h Handler) TestDelete(context *fiber.Ctx) error {
	context.Send([]byte("ok"))
	return nil
}
