package route

import (
	"github.com/labstack/echo/v4"
)

type handler interface {
	HandleGetUser(c echo.Context) error
	HandleGetUserById(c echo.Context) error
	HandleCreateUser(c echo.Context) error
}

func Register(e *echo.Echo, h handler) {
	e.GET("/user", h.HandleGetUser)
	e.GET("/user/:id", h.HandleGetUserById)
	e.POST("/user", h.HandleCreateUser)
}
