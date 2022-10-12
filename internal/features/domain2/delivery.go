package domain2

import "github.com/gofiber/fiber/v2"

const (
	FeatureName = "domain2"
)

type ApiHandler interface {
	TestGet(*fiber.Ctx) error
	TestPost(*fiber.Ctx) error
	TestPatch(*fiber.Ctx) error
	TestDelete(*fiber.Ctx) error
}
