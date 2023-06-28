package adapters

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func EchoHandlerAdapter(handlerFunc http.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		handlerFunc(c.Response().Writer, c.Request())

		return nil
	}
}
