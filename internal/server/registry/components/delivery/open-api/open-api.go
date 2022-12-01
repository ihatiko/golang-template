package open_api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

const (
	v1      = "/grpc/v1"
	metrics = "/metrics"
	health  = "/health"
)

type openApiContainer struct {
	V1    fiber.Router
	App   *fiber.App
	Debug bool
}

func NewOpenApiContainer(
	app *fiber.App,
	debug bool,
) *openApiContainer {

	return &openApiContainer{
		App:   app,
		V1:    app.Group(v1),
		Debug: debug,
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
