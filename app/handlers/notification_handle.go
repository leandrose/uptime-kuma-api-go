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

func NotificationsGetHandle(c echo.Context) error {
	var service uptimekuma.IUptimeKumaService
	err := container.Resolve(&service)
	if err != nil {
		return http.Error500Presenter(c.Response(), errors.New("failed instance service IUptimeKumaService"))
	}

	notifications := service.GetNotifications()
	return http.NotificationsPresenter(c.Response(), notifications)
}

func NotificationCreateHandle(c echo.Context) error {
	var service uptimekuma.IUptimeKumaService
	err := container.Resolve(&service)
	if err != nil {
		return http.Error500Presenter(c.Response(), errors.New("failed instance service IUptimeKumaService"))
	}

	notification := entities.Notification{}
	if err = c.Bind(&notification); err != nil {
		return http.Error400Presenter(c.Response(), err)
	}

	n, err := service.CreateNotification(notification)
	if err != nil {
		return http.Error500Presenter(c.Response(), err)
	}
	return http.NotificationPresenter(c.Response(), *n)
}

func NotificationDeleteHandle(c echo.Context) error {
	var service uptimekuma.IUptimeKumaService
	err := container.Resolve(&service)
	if err != nil {
		return http.Error500Presenter(c.Response(), errors.New("failed instance service IUptimeKumaService"))
	}

	notificationID, err := strconv.Atoi(c.Param("notification_id"))
	if err != nil {
		return http.Error400Presenter(c.Response(), err)
	}

	err = service.DeleteNotification(notificationID)
	if err != nil {
		return http.Error500Presenter(c.Response(), err)
	}
	return http.SuccessPresenter(c.Response())
}
