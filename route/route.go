package route

import (
	"github.com/labstack/echo/v4"
)

type handler interface {
	HandleLogin(c echo.Context) error
	HandleGetUser(c echo.Context) error
	HandleCreateUser(c echo.Context) error
	HandleTransact(c echo.Context) error
}

func Register(e *echo.Echo, h handler) {
	e.POST("/login", h.HandleLogin)
	e.GET("/user", h.HandleGetUser)
	e.POST("/user", h.HandleCreateUser)
	e.PUT("/transact", h.HandleTransact)
}
