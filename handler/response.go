package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func sendSuccessResponse(c echo.Context, data interface{}) error {
	return sendJsonResponse(c, http.StatusOK, "Ok", data)
}

func sendCreatedResponse(c echo.Context, data interface{}) error {
	return sendJsonResponse(c, http.StatusCreated, "Created", data)
}

func sendBadRequestResponse(c echo.Context, data interface{}, message string) error {
	return sendJsonResponse(c, http.StatusBadRequest, message, data)
}

func sendInternalErrorResponse(c echo.Context, data interface{}, message string) error {
	return sendJsonResponse(c, http.StatusInternalServerError, message, data)
}

func sendJsonResponse(c echo.Context, status int, message string, data interface{}) error {
	return c.JSON(status, map[string]interface{}{
		"message": message,
		"data":    data,
	})
}
