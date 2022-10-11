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

func (cnt *openApiContainer) Log(ctx *fiber.Ctx) error {
	t := time.Now()
	URL := ctx.OriginalURL()
	Method := ctx.Method()
	Status := ctx.Response().StatusCode()
	err := ctx.Next()
	Duration := time.Since(t)
	if cnt.Debug {
		log.HttpMiddlewareAccessLoggerDebug(Method, URL, Status, Duration, string(ctx.Request().Body()), string(ctx.Response().Body()))
	} else {
		log.HttpMiddlewareAccessLogger(Method, URL, Status, Duration)
	}
	return err
}
