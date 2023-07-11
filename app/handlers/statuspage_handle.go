package handlers

import (
	"errors"
	"github.com/golobby/container/v3"
	"github.com/labstack/echo/v4"
	"github.com/leandrose/uptime-kuma-api-go/domain/entities"
	"github.com/leandrose/uptime-kuma-api-go/domain/services/uptimekuma"
	"github.com/leandrose/uptime-kuma-api-go/presenter/http"
)

type statusPageUpdateDTO struct {
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Icon        *string `json:"icon"`
	//Theme             *string   `json:"theme"`
	//Published         *bool     `json:"published"`
	//ShowTags          *bool     `json:"showTags"`
	//DomainNameList    *[]string `json:"domainNameList"`
	//CustomCSS         *string   `json:"customCSS"`
	//FooterText        *string   `json:"footerText"`
	//ShowPoweredBy     *bool     `json:"showPoweredBy"`
	//GoogleAnalyticsId *string   `json:"googleAnalyticsId"`
}

func (dto *statusPageUpdateDTO) Init() {
	sp := entities.StatusPage{}
	sp.Init()

	dto.Title = sp.Title
	dto.Description = sp.Description
	dto.Icon = sp.Icon
	//dto.Theme = sp.Theme
	//dto.Published = sp.Published
	//dto.ShowTags = sp.ShowTags
	//dto.DomainNameList = sp.DomainNameList
	//dto.CustomCSS = sp.CustomCSS
	//dto.FooterText = sp.FooterText
	//dto.ShowPoweredBy = sp.ShowPoweredBy
	//dto.GoogleAnalyticsId = sp.GoogleAnalyticsId
}

type statusPageParamSlugDTO struct {
	Slug string `param:"slug"`
}

func StatusPagesGetHandle(c echo.Context) error {
	var service uptimekuma.IUptimeKumaService
	err := container.Resolve(&service)
	if err != nil {
		return http.Error500Presenter(c.Response(), errors.New("failed instance service IUptimeKumaService"))
	}

	return http.StatusPagesPresenter(c.Response(), service.GetStatusPages())
}

func StatusPageGetHandle(c echo.Context) error {
	slug := c.Param("slug")

	var service uptimekuma.IUptimeKumaService
	err := container.Resolve(&service)
	if err != nil {
		return http.Error500Presenter(c.Response(), errors.New("failed instance service IUptimeKumaService"))
	}

	statuspage, err := service.GetStatusPage(slug)
	if err != nil {
		return http.Error500Presenter(c.Response(), err)
	}

	return http.StatusPagePresenter(c.Response(), *statuspage)
}

func StatusPagesCreateHandle(c echo.Context) error {
	dto := struct {
		Slug  string `json:"slug" form:"slug"`
		Title string `json:"title" form:"title"`
	}{}
	err := c.Bind(&dto)
	if err != nil {
		return http.Error400Presenter(c.Response(), err)
	}

	statuspage := entities.StatusPage{}
	statuspage.Slug = dto.Slug
	statuspage.Title = dto.Title

	var service uptimekuma.IUptimeKumaService
	err = container.Resolve(&service)
	if err != nil {
		return http.Error500Presenter(c.Response(), errors.New("failed instance service IUptimeKumaService"))
	}

	ss, err := service.CreateStatusPage(statuspage)
	if err != nil {
		return http.Error500Presenter(c.Response(), err)
	}

	return http.StatusPagePresenter(c.Response(), *ss)
}

func StatusPageUpdateHandle(c echo.Context) error {
	dto := struct {
		Slug              string                     `param:"slug"`
		Title             string                     `json:"title"`
		Description       *string                    `json:"description"`
		Icon              *string                    `json:"icon"`
		Theme             *string                    `json:"theme"`
		Published         *bool                      `json:"published"`
		ShowTags          *bool                      `json:"showTags"`
		DomainNameList    *[]string                  `json:"domainNameList"`
		CustomCSS         *string                    `json:"customCSS"`
		FooterText        *string                    `json:"footerText"`
		ShowPoweredBy     *bool                      `json:"showPoweredBy"`
		GoogleAnalyticsId *string                    `json:"googleAnalyticsId"`
		PublicGroupList   []entities.PublicGroupList `json:"monitors"`
	}{}
	sp := entities.StatusPage{}
	sp.Init()
	dto.Title = sp.Title
	dto.Description = sp.Description
	dto.Icon = sp.Icon
	dto.Theme = sp.Theme
	dto.Published = sp.Published
	dto.ShowTags = sp.ShowTags
	dto.DomainNameList = sp.DomainNameList
	dto.CustomCSS = sp.CustomCSS
	dto.FooterText = sp.FooterText
	dto.ShowPoweredBy = sp.ShowPoweredBy
	dto.GoogleAnalyticsId = sp.GoogleAnalyticsId
	if err := c.Bind(&dto); err != nil {
		return http.Error400Presenter(c.Response(), err)
	}

	var service uptimekuma.IUptimeKumaService
	if err := container.Resolve(&service); err != nil {
		return http.Error500Presenter(c.Response(), errors.New("failed instance service IUptimeKumaService"))
	}

	statuspage, err := service.GetStatusPage(dto.Slug)
	if err != nil {
		return http.Error404Presenter(c.Response(), err)
	}
	statuspage.Title = dto.Title
	statuspage.Description = dto.Description
	statuspage.Icon = dto.Icon
	statuspage.Theme = dto.Theme
	statuspage.Published = dto.Published
	statuspage.ShowTags = dto.ShowTags
	statuspage.DomainNameList = dto.DomainNameList
	statuspage.CustomCSS = dto.CustomCSS
	statuspage.FooterText = dto.FooterText
	statuspage.ShowPoweredBy = dto.ShowPoweredBy
	statuspage.GoogleAnalyticsId = dto.GoogleAnalyticsId

	sp2, err := service.UpdateStatusPage(dto.Slug, *statuspage, dto.PublicGroupList)
	if err != nil {
		return http.Error500Presenter(c.Response(), err)
	}

	return http.StatusPagePresenter(c.Response(), *sp2)
}

func StatusPageDeleteHandle(c echo.Context) error {
	paramDto := struct {
		Slug string `param:"slug"`
	}{}
	if err := c.Bind(&paramDto); err != nil {
		return http.Error400Presenter(c.Response(), err)
	}

	var service uptimekuma.IUptimeKumaService
	if err := container.Resolve(&service); err != nil {
		return http.Error500Presenter(c.Response(), errors.New("failed instance service IUptimeKumaService"))
	}

	if err := service.DeleteStatusPage(paramDto.Slug); err != nil {
		return http.Error500Presenter(c.Response(), err)
	}

	return http.SuccessPresenter(c.Response())
}

func StatusPageCreateIncidentHandle(c echo.Context) error {
	paramDto := struct {
		Slug    string `param:"slug"`
		Style   string `json:"style"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}{}
	if err := c.Bind(&paramDto); err != nil {
		return http.Error400Presenter(c.Response(), err)
	}

	var service uptimekuma.IUptimeKumaService
	if err := container.Resolve(&service); err != nil {
		return http.Error500Presenter(c.Response(), errors.New("failed instance service IUptimeKumaService"))
	}

	incident := entities.StatusPageIncident{
		Style:   paramDto.Style,
		Title:   paramDto.Title,
		Content: paramDto.Content,
	}
	if _, err := service.PinIncidentStatusPage(paramDto.Slug, incident); err != nil {
		return http.Error500Presenter(c.Response(), err)
	}

	return http.SuccessPresenter(c.Response())
}

func StatusPageRemoveIncidentHandle(c echo.Context) error {
	paramDto := struct {
		Slug string `param:"slug"`
	}{}
	if err := c.Bind(&paramDto); err != nil {
		return http.Error400Presenter(c.Response(), err)
	}

	var service uptimekuma.IUptimeKumaService
	if err := container.Resolve(&service); err != nil {
		return http.Error500Presenter(c.Response(), errors.New("failed instance service IUptimeKumaService"))
	}

	if err := service.UnpinIncidentStatusPage(paramDto.Slug); err != nil {
		return http.Error500Presenter(c.Response(), err)
	}

	return http.SuccessPresenter(c.Response())
}
