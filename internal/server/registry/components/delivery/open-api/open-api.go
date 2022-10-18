package open_api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ihatiko/di"
	"test/internal/features/basket"
	"test/internal/features/payments"
	"test/internal/features/products"
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
	di.Invoke(func(h products.ApiHandler) {
		cnt.V1.Get(products.FeatureName, h.TestGet)
		cnt.V1.Post(products.FeatureName, h.TestPost)
		cnt.V1.Patch(products.FeatureName, h.TestPatch)
		cnt.V1.Delete(products.FeatureName, h.TestDelete)
	})

	di.Invoke(func(h payments.ApiHandler) {
		cnt.V1.Get(payments.FeatureName, h.TestGet)
		cnt.V1.Post(payments.FeatureName, h.TestPost)
		cnt.V1.Patch(payments.FeatureName, h.TestPatch)
		cnt.V1.Delete(payments.FeatureName, h.TestDelete)
	})

	di.Invoke(func(h basket.ApiHandler) {
		cnt.V1.Get(basket.FeatureName, h.TestGet)
		cnt.V1.Post(basket.FeatureName, h.TestPost)
		cnt.V1.Patch(basket.FeatureName, h.TestPatch)
		cnt.V1.Delete(basket.FeatureName, h.TestDelete)
	})
}
