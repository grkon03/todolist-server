package router

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetRoute(e *echo.Echo) error {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	api := e.Group("/api/v1")

	api.GET("/ping", Ping)

	return nil
}

func Ping(ctx echo.Context) error {
	log.Println("ping recieved")
	return ctx.String(http.StatusOK, "pong")
}
