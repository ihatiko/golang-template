package open_api

import (
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func (cnt *openApiContainer) Middlewares() {
	cnt.App.Use(favicon.New())
	cnt.App.Use(pprof.New())
	cnt.App.Use(requestid.New())
}
