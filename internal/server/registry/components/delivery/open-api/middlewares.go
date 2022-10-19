package open_api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/ihatiko/log"
	"time"
)

func (cnt *openApiContainer) Middlewares() {
	cnt.App.Use(cnt.Log)
	cnt.App.Use(favicon.New())
	cnt.App.Use(pprof.New())
	cnt.App.Use(requestid.New())
}

func (cnt *openApiContainer) Log(context *fiber.Ctx) error {
	t := time.Now()
	URL := context.OriginalURL()
	if URL == favIcon || URL == health || URL == metrics {
		return context.Next()
	}
	Method := context.Method()
	Status := context.Response().StatusCode()
	err := context.Next()
	Duration := time.Since(t)
	if cnt.Debug {
		log.HttpMiddlewareAccessLoggerDebug(Method, URL, Status, Duration, string(context.Request().Body()), string(context.Response().Body()))
	} else {
		log.HttpMiddlewareAccessLogger(Method, URL, Status, Duration)
	}
	return err
}
