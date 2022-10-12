package open_api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ihatiko/di"
	"test/internal/features/domain1"
	"test/internal/features/domain2"
	"test/internal/features/domain3"
)

const (
	v1      = "/api/v1"
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

func (cnt *openApiContainer) OpenApiRegistryV1() {
	cnt.ConfigureTestDomainV1()
}

func (cnt *openApiContainer) ConfigureTestDomainV1() {
	di.Invoke(func(h domain1.ApiHandler) {
		cnt.V1.Get(domain1.FeatureName, h.TestGet)
		cnt.V1.Post(domain1.FeatureName, h.TestPost)
		cnt.V1.Patch(domain1.FeatureName, h.TestPatch)
		cnt.V1.Delete(domain1.FeatureName, h.TestDelete)
	})

	di.Invoke(func(h domain2.ApiHandler) {
		cnt.V1.Get(domain2.FeatureName, h.TestGet)
		cnt.V1.Post(domain2.FeatureName, h.TestPost)
		cnt.V1.Patch(domain2.FeatureName, h.TestPatch)
		cnt.V1.Delete(domain2.FeatureName, h.TestDelete)
	})

	di.Invoke(func(h domain3.ApiHandler) {
		cnt.V1.Get(domain3.FeatureName, h.TestGet)
		cnt.V1.Post(domain3.FeatureName, h.TestPost)
		cnt.V1.Patch(domain3.FeatureName, h.TestPatch)
		cnt.V1.Delete(domain3.FeatureName, h.TestDelete)
	})
}
