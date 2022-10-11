package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ihatiko/log"
	"test/internal/features/test-domain"
	"test/internal/features/test-domain/models"
)

type Handler struct {
	Service test_domain.Service
}

func NewApiHandler(service test_domain.Service) *Handler {
	return &Handler{Service: service}
}

func (h Handler) TestGet(ctx *fiber.Ctx) error {
	ctx.Send([]byte("ok"))
	return nil
}

func (h Handler) TestPost(ctx *fiber.Ctx) error {
	log.Info("POST")
	str := &models.TestDomainResponse{
		Field1: "1",
		Field2: "2",
		Field3: "3",
	}

	return ctx.JSON(str)
}

func (h Handler) TestPatch(ctx *fiber.Ctx) error {
	ctx.Send([]byte("ok"))
	return nil
}

func (h Handler) TestDelete(ctx *fiber.Ctx) error {
	ctx.Send([]byte("ok"))
	return nil
}
