package handlers

import (
	"errors"
	"github.com/golobby/container/v3"
	"github.com/labstack/echo/v4"
	"github.com/leandrose/uptime-kuma-api-go/domain/services/uptimekuma"
	"github.com/leandrose/uptime-kuma-api-go/presenter/http"
	"strconv"
)

func PingHandle(c echo.Context) error {
	var service uptimekuma.IUptimeKumaService
	err := container.Resolve(&service)
	if err != nil {
		return http.Error500Presenter(c.Response(), errors.New("failed instance service IUptimeKumaService"))
	}

	monitorID, err := strconv.Atoi(c.Param("monitor_id"))
	if err != nil && !(monitorID > 0) {
		return http.Error400Presenter(c.Response(), err)
	}

	ping, err := service.GetPing(monitorID)
	if err != nil {
		return http.Error404Presenter(c.Response(), err)
	}

	return http.PingPresenter(c.Response(), monitorID, *ping)
}

func PingsHandle(c echo.Context) error {
	var service uptimekuma.IUptimeKumaService
	err := container.Resolve(&service)
	if err != nil {
		return http.Error500Presenter(c.Response(), errors.New("failed instance service IUptimeKumaService"))
	}

	pings := service.GetPings()

	return http.PingsPresenter(c.Response(), pings)
}
