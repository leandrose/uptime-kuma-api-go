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

func TagGetHandle(c echo.Context) error {
	var service uptimekuma.IUptimeKumaService
	err := container.Resolve(&service)
	if err != nil {
		return http.Error500Presenter(c.Response(), errors.New("failed instance service IUptimeKumaService"))
	}

	tagID, err := strconv.Atoi(c.Param("tag_id"))
	if err != nil || !(tagID > 0) {
		return http.Error400Presenter(c.Response(), err)
	}

	tag, err := service.GetTag(tagID)
	if err != nil {
		return http.Error500Presenter(c.Response(), err)
	}

	return http.TagPresenter(c.Response(), *tag)
}

func TagsGetHandle(c echo.Context) error {
	var service uptimekuma.IUptimeKumaService
	err := container.Resolve(&service)
	if err != nil {
		return http.Error500Presenter(c.Response(), errors.New("failed instance service IUptimeKumaService"))
	}

	tags, err := service.GetTags()
	if err != nil {
		return http.Error500Presenter(c.Response(), err)
	}

	return http.TagsPresenter(c.Response(), *tags)
}

func TagCreateHandle(c echo.Context) error {
	var service uptimekuma.IUptimeKumaService
	err := container.Resolve(&service)
	if err != nil {
		return http.Error500Presenter(c.Response(), errors.New("failed instance service IUptimeKumaService"))
	}

	var tag entities.Tag
	err = c.Bind(&tag)
	if err != nil {
		return http.Error400Presenter(c.Response(), err)
	}

	t, err := service.CreateTag(tag)
	if err != nil {
		return http.Error500Presenter(c.Response(), err)
	}

	return http.TagPresenter(c.Response(), *t)
}

func TagDeleteHandle(c echo.Context) error {
	var service uptimekuma.IUptimeKumaService
	err := container.Resolve(&service)
	if err != nil {
		return http.Error500Presenter(c.Response(), errors.New("failed instance service IUptimeKumaService"))
	}

	tagID, err := strconv.Atoi(c.Param("tag_id"))
	if err != nil || !(tagID > 0) {
		return http.Error400Presenter(c.Response(), err)
	}

	err = service.DeleteTag(tagID)
	if err != nil {
		return http.Error500Presenter(c.Response(), err)
	}

	return http.SuccessPresenter(c.Response())
}

func TagUpdateHandle(c echo.Context) error {
	var service uptimekuma.IUptimeKumaService
	err := container.Resolve(&service)
	if err != nil {
		return http.Error500Presenter(c.Response(), errors.New("failed instance service IUptimeKumaService"))
	}

	tagID, err := strconv.Atoi(c.Param("tag_id"))
	if err != nil || !(tagID > 0) {
		return http.Error400Presenter(c.Response(), err)
	}
	tag := entities.Tag{
		ID: &tagID,
	}
	err = c.Bind(&tag)
	if err != nil {
		return http.Error400Presenter(c.Response(), err)
	}

	t, err := service.UpdateTag(tag)
	if err != nil {
		return http.Error500Presenter(c.Response(), err)
	}

	return http.TagPresenter(c.Response(), *t)
}
