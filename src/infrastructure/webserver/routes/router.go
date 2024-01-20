package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.elastic.co/apm/module/apmechov4"
	"net/http"
)

func RegisterPath() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(apmechov4.Middleware())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${status} ${time_rfc3339} [${method}] ${uri} ; remote_ip=${remote_ip} ; error=${error}\n",
	}))

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "Pong")
	})

	//init routes

	return e

}
