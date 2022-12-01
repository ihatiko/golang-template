package files

import "github.com/gofiber/fiber/v2"

type GrpcHandlers interface {
	SaveImage(ctx fiber.Ctx) error
}
