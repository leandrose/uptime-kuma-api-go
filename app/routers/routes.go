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
	// AVG PING
	e.GET("/pings_average", handlers.PingsAverageHandle)
	e.GET("/ping_average/:monitor_id", handlers.PingAverageHandle)
	// UPTIME
	e.GET("/uptimes", handlers.UptimesHandle)
	e.GET("/uptime/:monitor_id", handlers.UptimeHandle)
	// TAGS
	e.GET("/tags", handlers.TagsGetHandle)
	e.POST("/tags", handlers.TagCreateHandle)
	e.GET("/tags/:tag_id", handlers.TagGetHandle)
	e.DELETE("/tags/:tag_id", handlers.TagDeleteHandle)
	e.PUT("/tags/:tag_id", handlers.TagUpdateHandle)
}
