package http

import (
	"fmt"
	"github.com/leandrose/uptime-kuma-api-go/app/middlewares"
	"github.com/leandrose/uptime-kuma-api-go/app/routers"
	"github.com/labstack/echo/v4"
)

type Server struct {
	port int
}

func NewServer(port int) *Server {
	return &Server{port: port}
}

func (s *Server) Start() error {
	e := echo.New()

	middlewares.LoadMiddlewares(e)
	routers.LoadRouters(e)
	return e.Start(fmt.Sprintf(":%d", s.port))
}
