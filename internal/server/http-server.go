package server

import (
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"test/internal/server/registry/components"
)

func (s *Server) StartHttpServer() {
	app := fiber.New(fiber.Config{
		AppName:      s.Config.Server.Name,
		WriteTimeout: s.Config.Server.WriteTimeout,
		ReadTimeout:  s.Config.Server.ReadTimeout,
		JSONDecoder:  jsoniter.Unmarshal,
		JSONEncoder:  jsoniter.Marshal,
	})
	container := components.NewOpenApiContainer(app)
	container.Middlewares()
	container.OpenApiRegistryV1()
	go app.Listen(s.Config.Server.Port)
}
