package handlers

import (
	"errors"
	"github.com/golobby/container/v3"
	"github.com/labstack/echo/v4"
	"github.com/leandrose/uptime-kuma-api-go/domain/services/uptimekuma"
	"github.com/leandrose/uptime-kuma-api-go/presenter/http"
	"strconv"
)

func UptimeHandle(c echo.Context) error {
	var service uptimekuma.IUptimeKumaService
	err := container.Resolve(&service)
	if err != nil {
		return http.Error500Presenter(c.Response(), errors.New("failed instance service IUptimeKumaService"))
	}

	monitorID, err := strconv.Atoi(c.Param("monitor_id"))
	if err != nil {
		http.Error400Presenter(c.Response(), err)
	}

	uptimes, err := service.GetUptime(monitorID)

	return http.UptimesPresenter(c.Response(), *uptimes)
}

func UptimesHandle(c echo.Context) error {
	var service uptimekuma.IUptimeKumaService
	err := container.Resolve(&service)
	if err != nil {
		return http.Error500Presenter(c.Response(), errors.New("failed instance service IUptimeKumaService"))
	}

	uptimes := service.GetUptimes()

	return http.UptimesPresenter(c.Response(), uptimes)
}
