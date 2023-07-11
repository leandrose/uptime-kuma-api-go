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
	em.PATCH("/resume", handlers.MonitorResumeHandle)
	em.PATCH("/pause", handlers.MonitorPauseHandle)
	em.GET("/heartbeats", handlers.MonitorHeartbeatsGetHandle)
	emt := em.Group("/tag/:tag_id")
	emt.POST("", handlers.MonitorTagAddHandle)
	emt.DELETE("", handlers.MonitorTagDeleteHandle)
	// NOTIFICATIONS
	e.GET("/notifications", handlers.NotificationsGetHandle)
	e.POST("/notifications", handlers.NotificationCreateHandle)
	e.PUT("/notifications", handlers.NotificationCreateHandle)
	e.DELETE("/notifications/:notification_id", handlers.NotificationDeleteHandle)
	// PING AVERAGE
	e.GET("/pings_average", handlers.PingsAverageHandle)
	e.GET("/ping_average/:monitor_id", handlers.PingAverageHandle)
	// UPTIME
	e.GET("/uptimes", handlers.UptimesHandle)
	e.GET("/uptime/:monitor_id", handlers.UptimeHandle)
	// STATUS PAGES
	es := e.Group("/statuspages")
	es.GET("", handlers.StatusPagesGetHandle)
	es.POST("", handlers.StatusPagesCreateHandle)
	ess := es.Group("/:slug")
	ess.GET("", handlers.StatusPageGetHandle)
	ess.POST("", handlers.StatusPageUpdateHandle)
	ess.DELETE("", handlers.StatusPageDeleteHandle)
	// STATUS PAGE INCIDENT
	ess.POST("/incident", handlers.StatusPageCreateIncidentHandle)
	ess.DELETE("/incident", handlers.StatusPageRemoveIncidentHandle)
	// TAGS
	e.GET("/tags", handlers.TagsGetHandle)
	e.POST("/tags", handlers.TagCreateHandle)
	e.GET("/tags/:tag_id", handlers.TagGetHandle)
	e.DELETE("/tags/:tag_id", handlers.TagDeleteHandle)
	e.PUT("/tags/:tag_id", handlers.TagUpdateHandle)
}
