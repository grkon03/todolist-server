package router

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (h *SqlHandler) Ping(ctx echo.Context) error {
	log.Println("ping recieved")

	pong := "pong"

	rand.Seed(time.Now().Unix())
	if rand.Intn(10) == 9 {
		pong = "ping"
	}

	return ctx.String(http.StatusOK, pong)
}
