package components

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ihatiko/fiber-logger-middleware"
	"test/internal/features/test-domain"
	domainApi "test/internal/features/test-domain/delivery/api"
	"test/internal/server/registry/components/feature-components/service"
)

const (
	v1 = "/api/v1"
)

type openApiContainer struct {
	V1  fiber.Router
	App *fiber.App
}

func NewOpenApiContainer(app *fiber.App) *openApiContainer {
	return &openApiContainer{
		App: app,
		V1:  app.Group(v1),
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

func (cnt *openApiContainer) Middlewares() {
	cnt.App.Use(fiber_logger_middleware.Log)
}
