package open_api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ihatiko/di"
	"test/internal/features/test-domain"
)

const (
	v1 = "/api/v1"
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

func (cnt *openApiContainer) OpenApiRegistryV1() {
	cnt.ConfigureTestDomainV1()
}

func (cnt *openApiContainer) ConfigureTestDomainV1() {
	di.Invoke(func(h test_domain.ApiHandler) {
		cnt.V1.Get(test_domain.FeatureName, h.TestGet)
		cnt.V1.Post(test_domain.FeatureName, h.TestPost)
		cnt.V1.Patch(test_domain.FeatureName, h.TestPatch)
		cnt.V1.Delete(test_domain.FeatureName, h.TestDelete)
	})
}
