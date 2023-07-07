package handlers

import (
	"errors"
	"github.com/golobby/container/v3"
	"github.com/labstack/echo/v4"
	"github.com/leandrose/uptime-kuma-api-go/domain/entities"
	"github.com/leandrose/uptime-kuma-api-go/domain/services/uptimekuma"
	"github.com/leandrose/uptime-kuma-api-go/presenter/http"
	"strconv"
)

func MonitorsHandle(c echo.Context) error {
	var service uptimekuma.IUptimeKumaService
	err := container.Resolve(&service)
	if err != nil {
		return http.Error500Presenter(c.Response(), errors.New("failed instance service IUptimeKumaService"))
	}

	return http.MonitorsPresenter(c.Response(), service.GetMonitors())
}

func MonitorGetByIdHandle(c echo.Context) error {
	monitorId, err := strconv.Atoi(c.Param("monitor_id"))
	if err != nil {
		return http.Error400Presenter(c.Response(), err)
	}
	var service uptimekuma.IUptimeKumaService
	err = container.Resolve(&service)
	if err != nil {
		return http.Error500Presenter(c.Response(), errors.New("failed instance service IUptimeKumaService"))
	}

	monitor, err := service.GetMonitorById(monitorId)
	if err != nil {
		return http.Error404Presenter(c.Response(), err)
	}

	return http.MonitorPresenter(c.Response(), *monitor)
}

func MonitorCreateHandle(c echo.Context) error {
	monitor := entities.NewMonitor()
	err := c.Bind(&monitor)
	if err != nil {
		return http.Error500Presenter(c.Response(), err)
	}

	var service uptimekuma.IUptimeKumaService
	err = container.Resolve(&service)
	if err != nil {
		return http.Error500Presenter(c.Response(), err)
	}

	m, err := service.CreateMonitor(monitor)
	if err != nil {
		return http.Error400Presenter(c.Response(), err)
	}
	return http.MonitorPresenter(c.Response(), *m)
}

func MonitorEditHandle(c echo.Context) error {
	monitorId, err := strconv.Atoi(c.Param("monitor_id"))
	if err != nil {
		return http.Error400Presenter(c.Response(), err)
	}

	var service uptimekuma.IUptimeKumaService
	err = container.Resolve(&service)
	if err != nil {
		return http.Error500Presenter(c.Response(), errors.New("failed instance service IUptimeKumaService"))
	}

	monitor, err := service.GetMonitorById(monitorId)
	if err != nil {
		return http.Error404Presenter(c.Response(), errors.New("monitor not found"))
	}
	err = c.Bind(monitor)
	if err != nil {
		return http.Error400Presenter(c.Response(), err)
	}

	m, err := service.CreateMonitor(*monitor)
	if err != nil {
		return http.Error500Presenter(c.Response(), err)
	}
	return http.MonitorPresenter(c.Response(), *m)
}

func MonitorDeleteHandle(c echo.Context) error {
	monitorId, err := strconv.Atoi(c.Param("monitor_id"))
	if err != nil {
		return http.Error400Presenter(c.Response(), err)
	}

	var service uptimekuma.IUptimeKumaService
	err = container.Resolve(&service)
	if err != nil {
		return http.Error500Presenter(c.Response(), errors.New("failed instance service IUptimeKumaService"))
	}

	err = service.DeleteMonitor(monitorId)
	if err != nil {
		return http.Error500Presenter(c.Response(), err)
	}

	return http.SuccessPresenter(c.Response())
}

func MonitorPauseHandle(c echo.Context) error {
	monitorId, err := strconv.Atoi(c.Param("monitor_id"))
	if err != nil {
		return http.Error400Presenter(c.Response(), err)
	}

	var service uptimekuma.IUptimeKumaService
	err = container.Resolve(&service)
	if err != nil {
		return http.Error500Presenter(c.Response(), errors.New("failed instance service IUptimeKumaService"))
	}

	err = service.PauseMonitor(monitorId)
	if err != nil {
		return http.Error500Presenter(c.Response(), err)
	}

	return http.SuccessPresenter(c.Response())
}

func MonitorResumeHandle(c echo.Context) error {
	monitorId, err := strconv.Atoi(c.Param("monitor_id"))
	if err != nil {
		return http.Error400Presenter(c.Response(), err)
	}

	var service uptimekuma.IUptimeKumaService
	err = container.Resolve(&service)
	if err != nil {
		return http.Error500Presenter(c.Response(), errors.New("failed instance service IUptimeKumaService"))
	}

	err = service.ResumeMonitor(monitorId)
	if err != nil {
		return http.Error500Presenter(c.Response(), err)
	}

	return http.SuccessPresenter(c.Response())
}
