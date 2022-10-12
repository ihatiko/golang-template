package domain1

import "github.com/gofiber/fiber/v2"

const (
	FeatureName = "domain1"
)

type ApiHandler interface {
	TestGet(*fiber.Ctx) error
	TestPost(*fiber.Ctx) error
	TestPatch(*fiber.Ctx) error
	TestDelete(*fiber.Ctx) error
}
