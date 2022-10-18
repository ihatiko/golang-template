package basket

import "github.com/gofiber/fiber/v2"

const (
	FeatureName = "payments"
)

type ApiHandler interface {
	TestGet(*fiber.Ctx) error
	TestPost(*fiber.Ctx) error
	TestPatch(*fiber.Ctx) error
	TestDelete(*fiber.Ctx) error
}
