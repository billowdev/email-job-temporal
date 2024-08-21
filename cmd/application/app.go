package application

import (
	"github.com/billowdev/email-job-temporal/internal/adapters/http/handlers"
	"github.com/billowdev/email-job-temporal/internal/adapters/http/routers"
	"github.com/billowdev/email-job-temporal/internal/core/services"
	"github.com/gofiber/fiber/v2"
	"go.temporal.io/sdk/client"
)

func AppContainer(app *fiber.App, temporalClient client.Client) *fiber.App {
	v1 := app.Group("/v1")
	route := routers.NewRoute(v1)
	EmailApp(route, temporalClient)
	return app
}

func EmailApp(r routers.RouterImpls, temporalClient client.Client) {
	emailSrv := services.NewEmailService(temporalClient)
	emailhandlers := handlers.NewEmailHandler(emailSrv)
	r.CreateEmailRoute(emailhandlers)
}
