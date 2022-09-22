package components

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/ihatiko/log"
	"test/internal/features/test-domain"
	domainApi "test/internal/features/test-domain/delivery/api"
	"test/internal/server/registry/components/feature-components/service"
	"time"
)

const (
	v1 = "/api/v1"
)

type openApiContainer struct {
	V1    fiber.Router
	App   *fiber.App
	Debug bool
}

func NewOpenApiContainer(app *fiber.App, debug bool) *openApiContainer {
	return &openApiContainer{
		App:   app,
		V1:    app.Group(v1),
		Debug: debug,
	}
}

func (cnt *openApiContainer) OpenApiRegistryV1() {
	container := service.NewContainer()
	cnt.ConfigureTestDomainV1(container)
}

func (cnt *openApiContainer) ConfigureTestDomainV1(container *service.Container) {
	handler := domainApi.NewApiHandler(container.Domain1Service)
	cnt.V1.Get(test_domain.FeatureName, handler.TestGet)
	cnt.V1.Post(test_domain.FeatureName, handler.TestPost)
	cnt.V1.Patch(test_domain.FeatureName, handler.TestPatch)
	cnt.V1.Delete(test_domain.FeatureName, handler.TestDelete)
}

func (cnt *openApiContainer) Log(ctx *fiber.Ctx) error {
	t := time.Now()
	URL := ctx.OriginalURL()
	Method := ctx.Method()
	Status := ctx.Response().StatusCode()
	ctx.Response().String()
	Duration := time.Since(t)
	headers := ctx.Request().Header
	headers = headers
	log.HttpMiddlewareAccessLogger(Method, URL, Status, 0, Duration)
	return nil
}

func (cnt *openApiContainer) Middlewares() {
	cnt.App.Use(cnt.Log)
	cnt.App.Use(favicon.New())
	cnt.App.Use(pprof.New())
	cnt.App.Use(requestid.New())
}
