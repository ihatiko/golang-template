package open_api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

const (
	metrics = "/metrics"
	health  = "/health"
)

type openApiContainer struct {
	App *fiber.App
}

func NewOpenApiContainer(
	app *fiber.App,
) *openApiContainer {

	return &openApiContainer{
		App: app,
	}
}

func (cnt *openApiContainer) ServicePoints() {
	cnt.App.Get(health, func(ctx *fiber.Ctx) error {
		ctx.Write([]byte("ok"))
		return nil
	})
	cnt.App.Get(metrics, func(ctx *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler())(ctx.Context())
		return nil
	})
}
