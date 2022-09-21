package api

import (
	"github.com/gofiber/fiber/v2"
	"test/internal/features/test-domain"
)

type Handler struct {
	Service test_domain.Service
}

func NewApiHandler(service test_domain.Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) TestGet(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) TestPost(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) TestPatch(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) TestDelete(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
