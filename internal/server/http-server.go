package server

import (
	"file_service/internal/server/registry/components/delivery/open-api"
	"github.com/gofiber/fiber/v2"
	"github.com/ihatiko/log"
	jsoniter "github.com/json-iterator/go"
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
	container := open_api.NewOpenApiContainer(app)

	container.Middlewares()
	container.ServicePoints()

	go func() {
		log.InfoF("Start http server %s", s.Config.Server.Port)
		err := app.Listen(s.Config.Server.Port)
		if err != nil {
			log.Fatal(err)
		}
	}()
	s.HttpServer = app
}
