package test_domain

import "github.com/gofiber/fiber/v2"

const (
	FeatureName = "test-domain"
)

type ApiHandler interface {
	TestGet(*fiber.Ctx) error
	TestPost(*fiber.Ctx) error
	TestPatch(*fiber.Ctx) error
	TestDelete(*fiber.Ctx) error
}
