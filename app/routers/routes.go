package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/leandrose/uptime-kuma-api-go/app/handlers"
	"github.com/leandrose/uptime-kuma-api-go/infra/http/adapters"
)

func LoadRouters(e *echo.Echo) {
	// add routes
	// recomendado criar diversos arquivos para melhor organizacao, caso houver muitas rotas

	e.GET("/hello-world", adapters.EchoHandlerAdapter(handlers.HelloWorldHandle))

	// MONITOR
	em := e.Group("/monitors")
	em.GET("", handlers.MonitorsHandle)
	em.POST("", handlers.MonitorCreateHandle)
	em = e.Group("/monitors/:monitor_id")
	em.GET("", handlers.MonitorGetByIdHandle)
	em.DELETE("", handlers.MonitorDeleteHandle)
	em.PATCH("", handlers.MonitorEditHandle)
}
