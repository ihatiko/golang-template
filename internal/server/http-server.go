package server

import (
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"test/internal/server/registry/components/delivery/open-api"
	"time"
)

func (s *Server) StartHttpServer() {
	app := fiber.New(fiber.Config{
		AppName:       s.Config.Server.Name,
		WriteTimeout:  time.Second * s.Config.Server.WriteTimeout,
		ReadTimeout:   time.Second * s.Config.Server.ReadTimeout,
		JSONDecoder:   jsoniter.Unmarshal,
		JSONEncoder:   jsoniter.Marshal,
		StrictRouting: true,
	})
	container := open_api.NewOpenApiContainer(app, s.Config.Server.Debug)

	container.Middlewares()
	container.ServicePoints()

	go app.Listen(s.Config.Server.Port)
	s.HttpServer = app
}
