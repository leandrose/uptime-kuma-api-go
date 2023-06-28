package routers

import (
	"github.com/leandrose/uptime-kuma-api-go/app/handlers"
	"github.com/leandrose/uptime-kuma-api-go/infra/http/adapters"
	"github.com/labstack/echo/v4"
)

func LoadRouters(e *echo.Echo) {
	// add routes
	// recomendado criar diversos arquivos para melhor organizacao, caso houver muitas rotas

	e.GET("/hello-world", adapters.EchoHandlerAdapter(handlers.HelloWorldHandle))
}
